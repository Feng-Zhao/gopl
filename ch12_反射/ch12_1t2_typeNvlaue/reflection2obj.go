package main

import (
	"fmt"
	"reflect"
)

// 对 reflect.Value 调用其 .Interface() 方法可以从 reflect.Value 转换为包含其 type和value 的 interface{} 类型
func main() {
	x := 3.4
	v := reflect.ValueOf(x)
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println("y = ", y)
	fmt.Println("y's type: ", reflect.TypeOf(y))
	fmt.Println("v = ", v)
	fmt.Println("v.Interface() = ", v.Interface())
	fmt.Println("v.Interface()'s type: ", reflect.TypeOf(v.Interface()))
}
