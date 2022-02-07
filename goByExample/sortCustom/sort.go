package main

import (
	"fmt"
	"sort"
)

// 自定义type
type byLength []string

// 实现 Len() Swap() Less() 接口
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// 调用 sort.Sort() 完成排序，内部会调用自定义的 Len(), Swap(), Less() 方法
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
