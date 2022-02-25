package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	change(s)
	fmt.Println(s)
}
func change(s []int) {
	s[0] = 0
	s[1] = 0
	s[2] = 0
}
