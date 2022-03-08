package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改 struct 内的字段
func main() {

	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	// 对 t 的指针取 elem
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("序号:%d: Name:%s Type:%s Value = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	// reset struct
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
