package main

import "fmt"

// 测试 defer 更改 return 的值
func main() {
	a := 1
	fmt.Println(add(a, 1))
	fmt.Println(add1(a, 1))
}

// 这种在函数体内部声明一个变量然后返回的,defer 无法更改返回值
func add(a int, b int) int {
	result := a
	defer func() {
		result = a + b
	}()
	return result
}

// 这种返回值提前声明的可以使用 defer 更改
func add1(a int, b int) (result int) {
	defer func() {
		result = a + b
	}()
	return result
}
