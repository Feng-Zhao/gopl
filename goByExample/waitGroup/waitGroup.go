package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// 创建一个 WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// 每启动一个 worker， WaitGroup + 1
		wg.Add(1)

		i := i

		go func() {
			// 每个worker结束后 WaitGroup.Done()
			// 注：如果在多个 func/goroutine 中传递 waitGroup, 应当使用指针
			defer wg.Done()
			worker(i)
		}()
	}

	// 这里会阻塞主线程，直到所有线程都结束， 类似于 countDownLatch
	wg.Wait()

}
