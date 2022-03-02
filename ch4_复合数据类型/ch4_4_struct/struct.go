package main

import "fmt"

type student struct {
	Name string
	id   int64
}

func main() {
	s := student{Name: "s", id: -1}
	fmt.Println(s)
	// 传递时是复制一份值,然后传递
	direct(s)
	fmt.Println(s) //{s -1}
	// 若想函数对结构体的更改影响外部变量,则需要传递指针
	pointer(&s)
	fmt.Println(s) //{sp 0}

}

func pointer(s *student) {
	s.Name = "sp"
	s.id = 0
}
func direct(s student) {
	s.Name = "sd"
	s.id = 1
}
