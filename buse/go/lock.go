package _go

import (
	"errors"
	"time"
)

type LockInfo struct {
	Key   string
	Token string
}

func Lock(key string, expireTime, delayWait int, isErr bool) (*LockInfo, error) {
	token := "token_1"
	closeCh := make(chan struct{})
	errCh := make(chan error, 1)

	go func() {
		for {
			select {
			case <-closeCh:
				return
			default:

			}

			if ok, err := getLock(key, token, expireTime, isErr); err != nil {
				errCh <- err
				return
			} else {
				if ok {
					errCh <- nil
					return
				}
			}
		}
	}()

	timer := time.After(time.Millisecond * time.Duration(delayWait))
	for {
		select {
		case <-timer:
			close(closeCh)
			return nil, errors.New("获取锁超时")
		case err := <-errCh:
			if err != nil {
				return nil, errors.New("获取锁错误")
			}

			return &LockInfo{
				Key:   key,
				Token: token,
			}, nil
		}
	}
}

func getLock(key, token string, expireTime int, isErr bool) (bool, error) {
	time.Sleep(time.Duration(expireTime) * time.Millisecond)
	if isErr {
		return false, errors.New("getLock错误")
	}
	return true, nil
}
