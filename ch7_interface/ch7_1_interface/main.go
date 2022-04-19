package main

import (
	"fmt"
	"log"

	"golang.org/x/net/html"
)

func main() {
	byteCount()

	log.Println("=========================")

	wordCount()

	log.Println("=========================")

	htmlDecoder("<html><head>lalala</head><body><p>hello world</p></body></html>")
}

func byteCount() {
	var c ByteCounter
	_, err := c.Write([]byte("hello"))
	if err != nil {
		panic(fmt.Sprintln("read failed", err))
	}
	fmt.Println(c) // 5

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12

}

// Ex 7.1 wordCounter
func wordCount() {
	var c WordCounter = 0
	fmt.Println(c) // 0
	_, err := c.Write([]byte("hello world"))
	if err != nil {
		panic(fmt.Sprintln("read failed", err))
	}
	fmt.Println(c) // 2
}

func htmlDecoder(s string) {
	reader := NewHtmlDecoder(s)

	z := html.NewTokenizer(reader)
	for z.Next() != html.ErrorToken {
		fmt.Printf("tagType: %s, Text: %s, Data: %s\n", z.Token().Type, z.Text(), z.Token().Data)

	}
	// node, err := html.Parse(reader)
	// if err != nil {
	// 	panic(err)
	// }
	// tag := node.FirstChild
	// for tag != nil {
	// 	fmt.Println(*tag)
	// 	tag = tag.NextSibling
	// }
}
