package main

import "fmt"

func main() {
	s := "hello world, 你好世界!"
	// len 返回字节数,而非字符数
	fmt.Printf("%d\n", len(s))
	// panic
	//p := s[len(s)]

	a := "123456789"
	// 可截取
	a1 := a[2:5]
	// 字符串不可变->不可更改, 所以都是重新给变量赋值
	//a1[0] = '0'
	//a1[1] = '0'
	//a1[2] = '0'
	fmt.Println(a1)
	// 原生字符串,即,按字面量解释的字符串,不理会转义字符\
	b := `\n\t\b\a\r\n`
	fmt.Printf("%s\n", b)
}
