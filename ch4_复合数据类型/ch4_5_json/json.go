package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	ID   int    `json:"student_id"`
	Name string `json:"student_name"`
}

// 使用json.Marshal() json.Unmarshal() 来对 go结构体和 json形式的 []byte 进行转换
func main() {
	s1 := student{ID: 1, Name: "s1"}
	jsonS1, err := json.Marshal(s1)
	// MarshalIndent() 序列化为带缩进,适合阅读的字符串,
	jsonS2, err := json.MarshalIndent(s1, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonS1)) // {"student_id":1,"student_name":"s1"}
	fmt.Println(string(jsonS2))
	/*
	   {
	           "student_id": 1,
	           "student_name": "s1"
	   }
	*/
	s2 := &student{}
	json.Unmarshal(jsonS1, s2)
	fmt.Println(*s2) // {1 s1}
}
