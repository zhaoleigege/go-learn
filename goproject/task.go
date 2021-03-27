package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	result, err := HandlerActivityFinishTask(context.Background())
	if err != nil {
		fmt.Printf("===========================%s=========\n", err)
		time.Sleep(1000 * time.Second)
		return
	}
	fmt.Printf("结果count: %d\n", result.TotalCount)
	time.Sleep(1000 * time.Second)
}

type TaskObj struct {
	Name string
	Time int
}

type TaskResult struct {
	TaskInfo
	Error error
}

type TaskInfo struct {
	TotalCount int
}

func HandlerActivityFinishTask(ctx context.Context) (*TaskInfo, error) {
	ctx, cancel := context.WithCancel(ctx)

	maxId := 100

	producer := make(chan []*TaskObj, 5)     // 接收从数据库读取的消息
	taskResultChan := make(chan *TaskResult) // 各个消费者执行的结果
	errChan := make(chan error)              // 错误处理
	// 生产者
	go func() {
		ProducerTask(ctx, maxId, errChan, producer)
	}()
	// 消费者
	go func() {
		consumerCount := 5
		wg := &sync.WaitGroup{}
		for i := 0; i < consumerCount; i++ {
			wg.Add(1)
			go func() {
				consumerTask(ctx, taskResultChan, producer, wg)
			}()

		}
		wg.Wait()
		close(taskResultChan)
		fmt.Println("消费者执行完成")
	}()

	// 统计结果
	totalCount := 0
	endChan := make(chan struct{})
	go func() {
		for taskResult := range taskResultChan {
			select {
			case <-ctx.Done():
				return
			default:

			}
			fmt.Println("还有消息接收")
			if taskResult.Error != nil {
				fmt.Printf("出现错误，统计结果执行完毕: %s\n", taskResult.Error)
				select {
				case <-ctx.Done():
					fmt.Printf("处理器强制退出: %s\n", taskResult.Error)
					return
				case errChan <- errors.New("执行消费者出错"):
					fmt.Printf("处理器自己退出: %s\n", taskResult.Error)
					return
				}
			}

			totalCount = totalCount + taskResult.TotalCount
		}
		fmt.Printf("处理器正常退出\n")
		close(endChan)
	}()

Loop:
	for {
		select {
		case err := <-errChan:
			if err != nil {
				cancel()
				// channel不一定需要手动关闭，当没有使用时GC会自动关闭
				// channel的关闭一定要在写入端而不是读取端
				//close(errChan)
				fmt.Println(err)
				return nil, errors.New("赠送积分消费错误")
			}

			break Loop
		case <-endChan:
			break Loop
		}
	}

	return &TaskInfo{
		TotalCount: totalCount,
	}, nil

}

// consumerTask 异步从协程里消费消息
func consumerTask(ctx context.Context, result chan<- *TaskResult, consumer <-chan []*TaskObj, wg *sync.WaitGroup) {
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

		taskInfo, err := handlerTickets(ctx, ticInfos)
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
					Error: errors.New("处理消息错误"),
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
func handlerTickets(ctx context.Context, ticInfos []*TaskObj) (*TaskInfo, error) {
	result := &TaskInfo{}
	if ticInfos[0].Time > 5 {
		return nil, errors.New("消费者错误")
	}

	result.TotalCount = len(ticInfos)
	return result, nil
}

// ProducerTask 生产可以赠送积分的订单列表
func ProducerTask(ctx context.Context, maxId int, errChan chan<- error, producer chan<- []*TaskObj) {
	if maxId <= 0 {
		return
	}

	gap := 10

	defer func() {
		fmt.Println("生产者生产完成")
		close(producer)
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
		totalTickets, err := getTaskObj(start, start+gap)
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
				case errChan <- errors.New("obj创建失败"):
					return
				}
			}
		}

		for _, tickets := range objSplit(totalTickets) {
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

func getTaskObj(startId, endId int) ([]*TaskObj, error) {
	objs := make([]*TaskObj, 0)
	for i := startId; i < endId; i++ {
		objs = append(objs, &TaskObj{
			Name: fmt.Sprintf("name-%d", i),
			Time: i,
		})
	}

	if startId > 10 {
		return nil, errors.New("生产者错误")
	}
	time.Sleep(1 * time.Second)
	return objs, nil
}

func objSplit(objs []*TaskObj) [][]*TaskObj {
	count := 5
	taskTwoArr := make([][]*TaskObj, 0)
	for i := 0; i < len(objs); i += count {
		endIndex := i + count

		if endIndex > len(objs) {
			endIndex = len(objs)
		}

		taskTwoArr = append(taskTwoArr, objs[i:endIndex])
	}

	return taskTwoArr
}
