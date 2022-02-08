package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	// struct1: {1 2}
	fmt.Printf("struct1: %v\n", p)
	// struct2: {x:1 y:2}
	fmt.Printf("struct2: %+v\n", p)
	// struct3: main.point{x:1, y:2}
	fmt.Printf("struct3: %#v\n", p)
	// type: main.point
	fmt.Printf("type: %T\n", p)
	// bool: true
	fmt.Printf("bool: %t\n", true)
	// int: 123
	fmt.Printf("int: %d\n", 123)
	// bin: 1110
	fmt.Printf("bin: %b\n", 14)
	// char: !
	fmt.Printf("char: %c\n", 33)
	// hex: 1c8
	fmt.Printf("hex: %x\n", 456)
	// float1: 78.900000
	fmt.Printf("float1: %f\n", 78.9)
	// float2: 1.234000e+08
	fmt.Printf("float2: %e\n", 123400000.0)
	// float3: 1.234000E+08
	fmt.Printf("float3: %E\n", 123400000.0)
	// str1: "string"
	fmt.Printf("str1: %s\n", "\"string\"")
	// str2: "\"string\""
	fmt.Printf("str2: %q\n", "\"string\"")
	// str3: 6865782074686973
	fmt.Printf("str3: %x\n", "hex this")
	// pointer: 0xc0000b2010
	fmt.Printf("pointer: %p\n", &p)
	// width1: |    12|   345|
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	// width2: |  1.20|  3.45|
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	// width3: |1.20  |3.45  |
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// width4: |   foo|     b|
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	// width5: |foo   |b     |
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// sprintf: a string
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	// io: an error
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
