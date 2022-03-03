package main

import (
	"fmt"
	"log"
	"time"
)

// 使用 defer + 闭包记录函数操作时长
func main() {
	defer trace("main")() //这里 trace 返回一个函数,这个函数是一个闭包,记录 start,然后 defer 调用这个函数,计算 time elapse
	time.Sleep(10 * time.Second)
}

func trace(name string) func() {
	start := time.Now()
	log.Printf("enter %s", name)
	return func() { fmt.Printf("exit %s, time elapsed %s seconds", name, time.Since(start)) }
}
