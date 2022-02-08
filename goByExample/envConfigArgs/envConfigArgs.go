package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 设置系统环境变量
	os.Setenv("FOO", "1")
	// 读取系统环境变量
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair)
	}
}
