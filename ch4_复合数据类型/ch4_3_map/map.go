package main

import "fmt"

func main() {
	// 新建 map 的两种种方式
	m := map[string]string{"a": "a", "b": "b", "c": "c"}
	var m2 = make(map[string]string)
	fillMap(m, "a")
	fmt.Println(m) // map[a:a b:a c:a] 函数内部更改 map,会更改外部的 map 内容

	if m2 == nil {
		fmt.Println("m2 is nil")
	} else {
		fmt.Println(m2) //map[]
	}

	var m3 map[string]string // map 的零值是 nil
	if m3 == nil {
		fmt.Println("m3 is nil") // m3 is nil
	} else {
		fmt.Println(m3)
	}
	// 向零值 map 中设置元素会引起 panic
	// m3["a"] = "a" //

	// 使用 delete 删除 map 中的元素
	delete(m, "a")
	fmt.Println(m) //map[b:a c:a]

	// map 中无对应值的键,会返回零值/空值,并且 ok 变量会返回 false
	v, ok := m["z"]
	fmt.Printf("v = %s, ok is %t \n", v, ok) //v = , ok is false

}
func fillMap(m map[string]string, c string) {
	for k, _ := range m {
		m[k] = c
	}
}
