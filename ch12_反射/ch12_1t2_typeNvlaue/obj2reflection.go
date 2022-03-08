package main

import (
	"fmt"
	"reflect"
)

// 使用
// reflect.TypeOf()
// reflect.ValueOf()
// 可以从 object 在运行时获取其 type 和 value
func main() {
	fmt.Println("================================")
	// reflect.TypeOf()
	// reflect.ValueOf()
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))   // type: float64
	fmt.Println("value:", reflect.ValueOf(x)) // value: 3.4
	// string() 方法深挖值, 而 fmt.Println() 会深入 reflect.Value 打印内部值
	fmt.Println("value:", reflect.ValueOf(x).String()) // value: <float64 Value>

	fmt.Println("================================")
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                               // type: float64
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) // kind is float64: true
	fmt.Println("value:", v.Float())                             // value: 3.4
	// fmt.Println("value:", v.Int())                               // panic, 无法转换, 使用时注意要先验证 Kind

	fmt.Println("================================")
	var y uint8 = 'y'
	v = reflect.ValueOf(y)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	y = uint8(v.Uint())                                       // v.Uint returns a uint64. 需要转换才能赋值
	// y = v.Uint()                                           // 无法通过编译

	fmt.Println("================================")
	type MyInt int
	var z MyInt = 7
	v = reflect.ValueOf(z)
	fmt.Println("kind is : ", v.Kind())          // kind is :  int  		// Kind() 无法分辨基础类型和自定义类型
	fmt.Println("Type is : ", reflect.TypeOf(z)) // Type is :  main.MyInt 	// reflect.TypeOf(z) 可以区分自定义类型
}
