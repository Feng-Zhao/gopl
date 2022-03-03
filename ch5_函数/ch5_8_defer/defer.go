package main

import "fmt"

func main() {
	fmt.Println("before defer")
	defer fmt.Println("defer processing")
	defer p()
	fmt.Println("after defer, before end")
}
func p() {
	fmt.Println("in func")
}
