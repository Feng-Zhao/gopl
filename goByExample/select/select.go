package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建 channel
	c1 := make(chan string)
	c2 := make(chan string)
	// 启动 协程
	go func() {
		time.Sleep(1 * time.Second)
		// 数据写入channel
		c1 <- "one"
	}()
	// 启动 协程
	go func() {
		time.Sleep(2 * time.Second)
		// 数据写入channel
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		// 数据从channel流出
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		// 数据从channel流出
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
