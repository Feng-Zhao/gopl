package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// 直接匹配 获取 bool 结果 结果： true
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	// 编译 正则表达式 供之后使用
	r, _ := regexp.Compile("p([a-z]+)ch")

	// true
	fmt.Println(r.MatchString("peach"))
	// peach
	fmt.Println(r.FindString("peach punch"))
	// idx: [0 5]
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	// [peach ea]  // ea 是 ([a-z]+) 的匹配
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// [0 5 1 3]
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// [peach punch pinch]
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	// [peach punch]
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// true
	fmt.Println(r.Match([]byte("peach")))
	// regexp: p([a-z]+)ch
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	// a <fruit>
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// a PEACH
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
