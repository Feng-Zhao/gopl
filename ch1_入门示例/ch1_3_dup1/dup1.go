package main

import (
	"bufio"
	"fmt"
	"os"
)

// 从stdin读取, 打印重复的行和重复次数
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 键不存在时 默认新建一个键值对,值为类型的空值
		counts[input.Text()]++
	}
	for line, n := range counts { // range 对 map 时, 返回 key, value 返回的顺序为随机的
		if n > 1 {
			// 格式化控制符 %d %s \t \n \b 等 详情见gopl书 p7 最下方
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
