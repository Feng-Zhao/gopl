package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 主要是 bufio.Scanner
	scanner := bufio.NewScanner(os.Stdin)
	// bufio.Scanner 读取一行
	for scanner.Scan() {
		// bufio.Scanner 返回之前 scan 的文本
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
