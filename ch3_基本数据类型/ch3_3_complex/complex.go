package main

import "fmt"

// complex 两种类型 complex64/complex128
// 内部是使用 float32/float64 实现
var x complex128 = complex(1, 2) // 复数 1+2i
//var y complex64 = complex(3, 4)
var y complex128 = complex(3, 4)

func main() {
	fmt.Println(x * y)
	fmt.Println(real(x * y)) // 实部
	fmt.Println(imag(x * y)) // 虚部
}
