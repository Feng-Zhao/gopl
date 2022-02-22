package main

import (
	"fmt"
	"os"
)

func main() {
	s, space := "", ""
	for _, arg := range os.Args[1:] { // range 对 array 和 slice 产生索引, 第一个值为 索引 i, 第二个值为 item[i]
		// 此处不需要第一个索引值 i, 可以以 _ (空标识符) 来接收该值,表示无用变量
		s += space + arg
		space = " "
	}
	fmt.Println(s)
}
