package main

import (
	"fmt"
	"time"
)

// timer 用于一次性延时
// ticker 多用于 重复性的定时动作
func main() {
	/**
	type Ticker struct {
		C <-chan Time // The channel on which the ticks are delivered.
		r runtimeTimer
	}
	*/
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			// 使用 channel 来控制 goroutine 的结束
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	// 结束 匿名 goroutine
	done <- true
	fmt.Println("Ticker stopped")
}
