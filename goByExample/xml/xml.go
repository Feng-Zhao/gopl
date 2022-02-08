package main

import (
	"encoding/xml"
	"fmt"
)

// xml 标签
type Plant struct {
	// 最外层
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// struct 转 xml 存储到 []byte
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// 设置 xml 头
	fmt.Println(xml.Header + string(out))

	// xml []byte 转 struct，使用指针接收
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	//
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	// 添加层级 使用 `>`
	type Nesting struct {
		// 最外层
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	// struct 转 xml
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
