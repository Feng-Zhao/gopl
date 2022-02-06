package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 使用 ioutil.ReadFile(filename) 一次性读取指定文件
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		// 使用 strings.Split() 分割字符串
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
