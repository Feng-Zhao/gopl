package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// counter
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// 使用 atomic 包下的方法 进行原子性操作
				atomic.AddUint64(&ops, 1)
				// 相对于 ops++ 会因为多线程操作而得到错误结果
				// ops++
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
