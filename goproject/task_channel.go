
type TaskResult struct {
	TaskInfo
	Error error
}

type TaskInfo struct {
	TotalCount        int
	SuccessTicCount   int
	FailAddCreditTics []int
	FailPushTics      []int
}

func HandlerActivityFinishTask(ctx context.Context, queryOnly bool) (*TaskInfo, error) {
	ctx, cancel := context.WithCancel(ctx)
	maxId, queryErr := projectkdbread.QueryMaxTicId(ctx)
	if queryErr != nil {
		return nil, errors.Errorf(queryErr, "获取最大id错误")
	}

	producer := make(chan []*projectkdbread.TicketInfoForCredit, 5) // 接收从数据库读取的消息
	taskResultChan := make(chan *TaskResult)                        // 各个消费者执行的结果
	errChan := make(chan error)                                     // 错误处理，这个并没有显示关闭，等待GC自动回收即可

	// 生产者
	go utils.HandlePanic(ctx, func() {
		ProducerTask(ctx, maxId, errChan, producer)
	})()

	// 消费者
	go utils.HandlePanic(ctx, func() {
		consumerCount := 5
		wg := &sync.WaitGroup{}

		for i := 0; i < consumerCount; i++ {
			wg.Add(1)
			go utils.HandlePanic(ctx, func() {
				consumerTask(ctx, taskResultChan, producer, wg, queryOnly)
			})()
		}
		wg.Wait()
		close(taskResultChan)
		logger.Infof("[%s]关闭消费者", utils.RequestIDFromContext(ctx))
	})()

	// 统计结果
	totalFacTics := make([]int, 0)
	totalFpTics := make([]int, 0)
	totalScount := 0
	totalCount := 0
	endChan := make(chan struct{}) // 表示处理器结果是否完成
	go utils.HandlePanic(ctx, func() {
		for taskResult := range taskResultChan {
			select {
			case <-ctx.Done():
				logger.Infof("[%s] 强制退出任务管理器", utils.RequestIDFromContext(ctx))
				return
			default:
			}

			if taskResult.Error != nil {
				select {
				case <-ctx.Done():
					logger.Infof("[%s] 强制退出任务管理器", utils.RequestIDFromContext(ctx))
					return
				case errChan <- errors.Errorf(taskResult.Error, "执行消费者出错"):
					logger.Infof("[%s] 消费者处理失败，退出任务管理器", utils.RequestIDFromContext(ctx))
					return
				}
			}

			totalCount = totalCount + taskResult.TotalCount
			totalFacTics = append(totalFacTics, taskResult.FailAddCreditTics...)
			totalFpTics = append(totalFpTics, taskResult.FailPushTics...)
			totalScount = totalScount + taskResult.SuccessTicCount
		}
		close(endChan)
		logger.Infof("[%s]关闭任务处理器", utils.RequestIDFromContext(ctx))
	})()

Loop:
	for {
		select {
		case <-endChan:
			break Loop
		case err := <-errChan:
			if err != nil {
				cancel()
				return nil, errors.Errorf(err, "赠送积分消费错误")
			}

			break Loop
		}
	}

	return &TaskInfo{
		TotalCount:        totalCount,
		SuccessTicCount:   totalScount,
		FailAddCreditTics: totalFacTics,
		FailPushTics:      totalFpTics,
	}, nil
}

// consumerTask 异步从协程里消费消息
func consumerTask(ctx context.Context, result chan<- *TaskResult, consumer <-chan []*projectkdbread.TicketInfoForCredit,
	wg *sync.WaitGroup, queryOnly bool) {
	defer func() {
		wg.Done()
	}()

	for ticInfos := range consumer {
		select {
		case <-ctx.Done():
			return
		default:
			break
		}

		taskInfo, err := handlerTickets(ctx, ticInfos, queryOnly)
		if err != nil {
			for {
				select {
				case <-ctx.Done():
					return
				default:

				}

				select {
				case <-ctx.Done():
					return
				case result <- &TaskResult{
					Error: errors.Errorf(err, "处理消息错误"),
				}:
					return
				}
			}
		}

		select {
		case <-ctx.Done():
			return
		default:
		}
		select {
		case <-ctx.Done():
			return
		case result <- &TaskResult{
			TaskInfo: *taskInfo,
		}:
			break
		}
	}
}

// handlerTickets 执行赠送积分以及发送邮件的函数
func handlerTickets(ctx context.Context, ticInfos []*projectkdbread.TicketInfoForCredit, queryOnly bool) (*TaskInfo, error) {
	requestId := utils.RequestIDFromContext(ctx)
	result := &TaskInfo{}
	result.TotalCount = len(ticInfos)
	result.FailAddCreditTics = make([]int, 0)
	result.FailPushTics = make([]int, 0)

	if queryOnly {
		return result, nil
	}

	for _, tickets := range projectkdbread.TicketInfoSplit(ticInfos) {
		if len(tickets) < 1 {
			continue
		}

		userIds := make([]int, 0, len(tickets))
		for _, ticket := range tickets {
			if !utils.IntIsIn(ticket.UserId, userIds...) {
				userIds = append(userIds, ticket.UserId)
			}
		}
		userMap := make(map[int64]*user.User, len(userIds))
		for _, ids := range ordercomm.DivideIntIDsSlice(userIds, 50) {
			if len(ids) <= 0 {
				break
			}

			tmpMap, err := user.GetUserMapByIds(ctx, ids)
			if err != nil {
				return nil, errors.Errorf(err, "获取用户信息错误")
			}

			for k, v := range tmpMap {
				userMap[k] = v
			}
		}

		for _, ticket := range tickets {
			if userInfo, ok := userMap[int64(ticket.UserId)]; ok {
				if userInfo.UserChannel == user_api.GC {
					err := projectkdb.SetIsCreditAdded(ctx, ticket.TicketID, CanNotSendCredit)
					if err != nil {
						err = errors.Errorf(err, "SetIsCreditAdded failed,ticketId:%d", ticket.TicketID)
						logger.Error(requestId, "[HandlerExecActivityFinishedTask] GC用户不需要赠送积分 set Failed err:", err)
						result.FailAddCreditTics = append(result.FailAddCreditTics, ticket.TicketID)
					}
					// guest checkout用户不赠送积分
					continue
				}
			}

			ordInfo, err := projectkdbread.QueryOrderByOrdId(ctx, ticket.OrderId)
			if err != nil {
				return nil, errors.Errorf(err, "通过orderId:%d 获取order信息出错", ticket.OrderId)
			}

			if ordInfo.PaymentChannel == orderconstpb.OrderPaymentChannelWechatGlobalPay || ordInfo.OrderChannel != OrderChannelNormal {
				err = projectkdb.SetIsCreditAdded(ctx, ticket.TicketID, CanNotSendCredit)
				if err != nil {
					err = errors.Errorf(err, "SetIsCreditAdded failed,ticketId:%d", ticket.TicketID)
					logger.Error(requestId, "[HandlerExecActivityFinishedTask]  不需要赠送积分 set Failed err:", err)
					result.FailAddCreditTics = append(result.FailAddCreditTics, ticket.TicketID)
				}
				// 不发送到队列
				continue
			}

			err = projectkdb.SetIsCreditAdded(ctx, ticket.TicketID, CanSendCredit)
			if err != nil {
				err = errors.Errorf(err, "SetIsCreditAdded failed,ticketId:%d", ticket.TicketID)
				logger.Error(requestId, "[HandlerExecActivityFinishedTask] Failed err:", err)
				result.FailAddCreditTics = append(result.FailAddCreditTics, ticket.TicketID)
				continue
			}

			ticket := &TicketInfo{
				TicketId:    ticket.TicketID,
				IsRepublish: false,
			}

			err = pub.Publish(requestId, pub.Type_FinishActivity, ticket)
			if err != nil {
				err = errors.Errorf(err, "Topic FinishActivity Publish failed,ticketId:%d", ticket.TicketId)
				logger.Error(requestId, "[HandlerExecActivityFinishedTask] Failed err:", err)
				result.FailPushTics = append(result.FailPushTics, ticket.TicketId)
				continue
			}
			result.SuccessTicCount++
		}
	}

	return result, nil
}

// ProducerTask 生产可以赠送积分的订单列表
func ProducerTask(ctx context.Context, maxId int, errChan chan<- error, producer chan<- []*projectkdbread.TicketInfoForCredit) {
	if maxId <= 0 {
		return
	}

	gap := 1000000

	defer func() {
		close(producer)
		logger.Infof("[%s]关闭生产者", utils.RequestIDFromContext(ctx))
	}()

	// 开始分段检索数据
	for start := 0; start < maxId; start += gap {
		select {
		case <-ctx.Done():
			return

		default:
			break
		}

		var err error
		totalTickets, err := projectkdbread.QueryAllCreditNotAddedTicketIDs(ctx, time.Now(), start, start+gap)
		if err != nil {
			for {
				select {
				case <-ctx.Done():
					return
				default:

				}

				select {
				case <-ctx.Done():
					return
				case errChan <- errors.Errorf(err, "获取需要赠送积分的ticket错误"):
					return
				}
			}
		}

		for _, tickets := range projectkdbread.TicketInfoSplit(totalTickets) {
			select {
			case <-ctx.Done():
				return
			default:
				producer <- tickets
				break
			}
		}
	}
}