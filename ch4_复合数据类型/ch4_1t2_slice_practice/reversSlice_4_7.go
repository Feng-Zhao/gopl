package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "你好,世界!"
	fmt.Println(str)
	fmt.Println(revers4_7([]byte(str)))
}
func revers4_7(arr []byte) string {
	rs := bytes.Runes(arr)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
