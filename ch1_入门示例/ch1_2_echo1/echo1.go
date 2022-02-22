package main

import (
	"fmt"
	"os"
)

func main() {
	var s, space string
	for i := 1; i < len(os.Args); i++ { // ++ 相当于一个语句 j = i++ 不合法,
		// 另外 go 中只有后缀形式的 ++,-- 没有前缀形式的 ++i, --i
		s += space + os.Args[i]
		space = " "
	}
	fmt.Println(s)
}
