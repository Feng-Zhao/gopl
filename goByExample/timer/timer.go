package main

import (
	"fmt"
	"time"
)

func main() {

	/** time.Timer
	type Timer struct {
		C <-chan Time
		r runtimeTimer
	}
	*/
	// 创建一个 time.Timer 结构如上
	timer1 := time.NewTimer(2 * time.Second)
	// C 是 timer 中的一个 channel， 下面一句会阻塞 main ggoroutine 直到 2s 后 timer 向 channel 写入
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// timer 相比于 time.Sleep, 优势是可以取消
	// time.Sleep(3 * time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
