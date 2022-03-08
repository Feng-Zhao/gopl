package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// v.SetFloat(7.1) // Error: will panic.

	var y float64 = 3.4
	v = reflect.ValueOf(y)
	fmt.Println("settability of v:", v.CanSet()) // 需要检验 reflec.Value 类型的 settability

	fmt.Println("===========================")
	/*
		简单来讲,一个变量是否可set,取决于是否地址可达,也就是是否持有原有变量的指针.
		想要通过反射来 set 一个变量,需要对其指针使用 reflect.ValueOf()
	*/
	var z float64 = 3.4
	p := reflect.ValueOf(&z)                     // Note: take the address of x.
	fmt.Println("type of p:", p.Type())          // type of p: *float64
	fmt.Println("settability of p:", p.CanSet()) // 这里反射对象依然不可设置,因为他是一个指针,但是我们要 set 的不是指针p, 而是 *p, 需要使用 Elem() 方法拿到指针所指的元素

	e := p.Elem()
	fmt.Println("settability of e:", e.CanSet())
	fmt.Println("before set, e = ", e)
	e.SetFloat(9.9)
	fmt.Println("after set, e = ", e)
}
