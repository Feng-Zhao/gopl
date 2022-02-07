package main

import (
	"fmt"
	"time"
)

// 通过 channel 阻塞线程，以确保线程接收请求的速率不会过高
func main() {
	// 模拟 5 个请求， 打到 channel 中
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 创建一个 Tick 结构为 一个 <-chan time.Time
	limiter := time.Tick(2000 * time.Millisecond)

	// 处理请求
	for req := range requests {
		// 使用 limite 阻塞线程
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("====== 令牌桶式 ======")
	// 创建带 buffer 的 channel， 以应对突发流量
	burstyLimiter := make(chan time.Time, 3)

	// 初始化 放入 3 个令牌
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 开启一个线程，每隔一定间隔时间 放入一块令牌
	go func() {
		for t := range time.Tick(2000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 创建请求 channel
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	// 处理求情
	for req := range burstyRequests {
		// 允许 瞬时流量 为 3， 即 buffer channel 的缓存大小
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
