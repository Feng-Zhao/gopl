package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123456"))
}

func comma(str string) string {
	if len(str) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	count := 0
	for i := len(str) - 1; i > 0; i-- {
		buf.WriteByte(str[i])
		count++
		if count%3 == 0 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(str[0])
	return revert(buf.String())
}
func revert(str string) string {
	buf := bytes.Buffer{}
	for i := len(str) - 1; i >= 0; i-- {
		buf.WriteByte(str[i])
	}
	return buf.String()
}
