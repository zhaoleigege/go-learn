package sync

import (
	"fmt"
	"sync"
)

// 参考 https://mp.weixin.qq.com/s/OcLrO-oINk2j2w9sEJvkPw
func StudentCond() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	var count uint64
	var stuSlice []int

	for i := 0; i < 30; i++ {
		go func(index int) {
			cond.L.Lock()
			stuSlice = append(stuSlice, index)
			count++
			cond.L.Unlock()

			cond.Broadcast()
		}(i)
	}

	cond.L.Lock()
	for count != 30 {
		cond.Wait()
	}
	cond.L.Unlock()

	fmt.Println(len(stuSlice), stuSlice)
}
