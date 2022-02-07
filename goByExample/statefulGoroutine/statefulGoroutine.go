package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int // resound channel
}
type writeOp struct {
	key  int
	val  int
	resp chan bool // resound channel
}

// 使用 goroutine + channel 完成同步操作， 相当于 wait notify
func main() {

	var readOps uint64
	var writeOps uint64

	// 读请求队列
	reads := make(chan readOp)
	// 写请求队列
	writes := make(chan writeOp)

	// 开启线程处理读写请求
	go func() {
		var state = make(map[int]int)
		for {
			select {
			// 读请求到来
			case read := <-reads:
				// 读取 并 释放读请求
				read.resp <- state[read.key]
				// 写请求到来
			case write := <-writes:
				// 写入
				state[write.key] = write.val
				// 释放写请求
				write.resp <- true
			}
		}
	}()

	// 开启 100 个线程
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// 创建 readOp
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				// 打入读请求队列 reads
				reads <- read
				// 阻塞线程
				<-read.resp
				// 读请求完成数++
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 开 10 个写线程
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// 创建写求情
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				// 打入写请求队列 writes
				writes <- write
				// 阻塞写请求
				<-write.resp
				// 写请求完成数++
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 主线程 睡眠1秒 等待协程工作
	time.Sleep(time.Second)

	// 获取最终的读写操作完成数量
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
