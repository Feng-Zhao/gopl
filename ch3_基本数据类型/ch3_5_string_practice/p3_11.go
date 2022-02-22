package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma2("+123456.123456"))
}

func comma2(str string) string {
	var flag = false
	var flagB = ""
	if strings.HasPrefix(str, "+") || strings.HasPrefix(str, "-") {
		flag = true
		flagB = str[0:1]
		str = str[1:]
	}
	if len(str) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	count := 0
	for i := len(str) - 1; i > 0; i-- {
		buf.WriteByte(str[i])
		if str[i] != '.' {
			count++
			if count%3 == 0 {
				buf.WriteByte(',')
			}
		} else {
			count = 0
		}

	}
	buf.WriteByte(str[0])
	result := revert2(buf.String())
	if flag {
		result = flagB + result
	}
	return result
}
func revert2(str string) string {
	buf := bytes.Buffer{}
	for i := len(str) - 1; i >= 0; i-- {
		buf.WriteByte(str[i])
		if str[i] == '.' {
			if i-1 >= 0 && str[i-1] == ',' {
				i--
			}
		}
	}
	return buf.String()
}
