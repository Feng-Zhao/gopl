package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

// 只有可导出的域，即开头为大写的成员变量才能配置json
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// type to json
	// 基础类型 使用 json.Marshal() 编码为 json 格式
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// slice 使用 json.Marshal() 编码为 json 格式
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	// map 使用 json.Marshal() 编码为 json 格式
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// struct json.Marshal() 编码为 json 格式
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// struct json.Marshal() 编码为 json 格式, 使用 json 标签制定编码为json时使用的域名 `json:"page"`
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// josn to type

	// json data 用 []byte 存储
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// 接收 json 的 map， 形式为 map[string]interface{}
	var dat map[string]interface{}

	// 将 json 转为 map, byt 为 json, 存储格式为[]byte， &dat 为接收方map
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	// 从 map 中取出值使用时需要强转类型
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// json to slice, 依然需要传指针作为接收方
	str := `["a","b","c"]`
	arr := []string{}
	if e := json.Unmarshal([]byte(str), &arr); e != nil {
		fmt.Println(e)
	}
	fmt.Println(arr)

	// json to struct, 直接用 struct 接收， 这里就不需要map形式接收时的类型转换了
	str = `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 创建 json 编码器，并指定输出流
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	// 将 map 转为 json 字符串
	enc.Encode(d)
}
