package main

import "fmt"

func main() {
	defer fmt.Println("outer")
	f()
}

func f() {
	fmt.Println("f.before")
	defer fmt.Println("inner")
	fmt.Println("f.after")
}
