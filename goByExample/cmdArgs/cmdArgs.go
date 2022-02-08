package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os.Args 获取 args
	// os.Args[0] 为程序名， 参数从 os.Args[1:] 开始
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
