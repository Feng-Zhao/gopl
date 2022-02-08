package main

import (
	"flag"
	"fmt"
)

func main() {
	// 设置 需要的 flag 的 （类型，默认值，提示语），flag 的方法返回的都是指针
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// 把指针与flag关联
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	// 初始化，相当于把输入的参数和程序内的变量绑定
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
