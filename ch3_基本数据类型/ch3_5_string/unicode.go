package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/*
 utf-8 编码
 0xxxxxxx 							 0-127 完全兼容 ASCII 码
 110xxxxx 10xxxxxx 					 两字节编码
 1110xxxx 10xxxxxx 10xxxxxx			 三字节编码
 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx 四字节编码
*/

func main() {
	s := "hello, 世界"
	fmt.Printf("字节长度: %d\n", len(s))
	fmt.Printf("rune长度: %d\n", utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		// 返回字符和 rune 长度
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// range 对字符串默认按 utf-8形式解码
	var count = 0
	for i, v := range s {
		fmt.Printf("%d \t%c\n", i, v)
		count++
	}
	fmt.Println("字符数量: " + strconv.Itoa(count))

}
