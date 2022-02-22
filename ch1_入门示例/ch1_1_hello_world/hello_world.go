package main // 主程序的启动包 应设置为 main

import "fmt"

func main() { // 主程序人口 固定为 main 函数
	fmt.Println("Hello 世界") // 默认 UTF-8 编码
}

// go run .\hello_world.go 编译并运行源码
// go build .\hello_world.go 将源码编译为可执行程序
