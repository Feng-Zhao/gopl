package main

import "fmt"

// 变量 声明
// var name type = value
var foo string = "foo"

// var name type 省略 value, 将初始化为类型的零值
var i int

// var name value,省略 type, 可以根据 value 自动推断合适的 type
var b = true

// short naming -> name := value
// only be used in func
// str := "hello"

// 常量声明
const zero = 0
const yes = true
const no = false

// type 声明
type myInt int
type myStruct struct {
	i   int
	b   bool
	str string
}

// 函数声明(主函数)
func main() {
	// short naming -> name := value
	// only be used in func
	str := "hello"
	fmt.Printf("%d\n%b\n%s\n", i, b, str)
}

// 函数声明
func add(a int, b int) int {
	fmt.Printf("the sum of a and b is: %d", a+b)
	return a + b
}
