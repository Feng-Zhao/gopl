package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[10]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
