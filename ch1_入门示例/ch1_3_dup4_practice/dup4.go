package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	files := os.Args[1:]
	// 未提供文件名,使用 stdin
	if len(files) == 0 {
		countLines(os.Stdin, counts, foundIn)
	} else {
		//提供文件名列表
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// 统计行出现次数,并记录文件名
			countLines(f, counts, foundIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, foundIn[line], line)
		}
	}
}

// 判断行文件名是否已经被记录
func in(needle string, strings []string) bool {
	for _, s := range strings {
		if needle == s {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)
	// 逐行扫描
	for input.Scan() {
		line := input.Text()
		// 记录行出现次数
		counts[line]++
		// 记录所属文件的文件名
		if !in(f.Name(), foundIn[line]) {
			foundIn[line] = append(foundIn[line], f.Name())
		}
	}
}
