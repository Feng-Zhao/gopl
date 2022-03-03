package main

import (
	"fmt"
	"time"
)

func main() {
	type needRecover struct{}

	defer func() {
		switch p := recover(); p {
		case nil: // 没有宕机
		case needRecover{}:
			fmt.Printf("发生宕机,且恢复\n")
		default:
			panic(p)
		}
	}()

	if t := time.Now(); t.Unix()%2 == 0 {
		panic(needRecover{})
	} else {
		panic(1)
	}
	fmt.Println("恢复之后还可以继续运行??") // 无法执行到这里
}
