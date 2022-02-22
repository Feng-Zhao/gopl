package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int to string
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println("=================")

	// string to int
	a, _ := strconv.Atoi("123")
	b, _ := strconv.ParseInt("ff", 16, 64)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("=================")

	// 对数字进行不同进制的格式化
	fmt.Println(strconv.FormatInt(0xff, 16))
	fmt.Println(strconv.FormatInt(0xff, 10))
	fmt.Printf("%d,%[1]b,%#[1]o,%#[1]x", x)

}
