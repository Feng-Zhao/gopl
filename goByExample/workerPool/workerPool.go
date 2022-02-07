// 使用 channel 分发 任务给 不同的 worker, 不必关心那个 worker 接收到了任务
// 相当于 线程池的概念
package main

import (
	"fmt"
	"time"
)

// 定义 worker
func worker(id int, jobs <-chan int, results chan<- int) {
	// 从 jobs channel 中读取任务
	for j := range jobs {
		// start
		fmt.Println("worker", id, "started  job", j)
		// 模拟计算过程
		time.Sleep(time.Second)
		// end
		fmt.Println("worker", id, "finished job", j)
		// 将结果 写入 result channel
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// 创建 3 个 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 将 5 个任务 放入 jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 从 result channel 中读取结果
	for a := 1; a <= numJobs; a++ {
		// 这里结果没什么用 直接丢掉
		<-results
	}
}
