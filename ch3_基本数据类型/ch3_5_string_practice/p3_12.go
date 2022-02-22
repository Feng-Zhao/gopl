package main

import "fmt"

func main() {
	fmt.Println(check())
}

func check() bool {
	s1 := "123456"
	s2 := "654321"
	m := make(map[rune]int)
	for _, v := range s1 {
		m[v] = m[v] + 1
	}
	m2 := make(map[rune]int)
	for _, v := range s2 {
		if m[v] == 0 {
			return false
		}
		m2[v] = m2[v] + 1
	}
	if len(m) != len(m2) {
		return false
	}
	return true
}
