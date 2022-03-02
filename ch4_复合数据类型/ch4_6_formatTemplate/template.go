package main

import (
	"fmt"
	"os"
	"text/template"
)

// 定义模板字符串
const temp = `{{.Name}} is a template, the value of a is {{.A}}, the value of b is {{.B}}
a+b is {{Add .A .B}}`

type Data struct {
	Name string
	A    int
	B    int
}

func main() {
	data := Data{Name: "data", A: 1, B: 2}
	// 创建模板
	result, err := template.New("report").Funcs(template.FuncMap{"Add": Add}).Parse(temp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 渲染模板
	result.Execute(os.Stdout, data)
}

func Add(a, b int) int {
	return a + b
}
