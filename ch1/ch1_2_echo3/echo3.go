package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	useStringsJoin()
	noFormat()
}

func useStringsJoin() {
	// 使用 strings.Join(elem []string, sep string) 来连接string数组中的内容,并且使用 sep指定的分隔符
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func noFormat() {
	// 直接打印数组, 会用 () 把数组内容括起来
	fmt.Println(os.Args[1:])
}
