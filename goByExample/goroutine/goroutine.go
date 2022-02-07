package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// time.Sleep(time.Second)
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		// 模拟任务准备时间
		// time.Sleep(time.Second)
		jobs <- j
		fmt.Println("sent job", j)
	}
	// 关闭 channel
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
