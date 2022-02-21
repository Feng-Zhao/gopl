package ch2_3_new

import "fmt"

// 使用 new(Type) 来创建变量,new()返回的为指针,即,*Type
func newType() {
	// new() 将值初始化为零值,并返回指针
	p := new(int)
	*p++
	fmt.Printf("%d", *p)
}
