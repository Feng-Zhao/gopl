package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fillZero(arr)
	fmt.Println(arr) //[1 2 3]

	fillZeroWithPointer(&arr)
	fmt.Println(arr) //[0 0 0]

	s := arr[:]
	fillZeroForSlice(s)
	fmt.Println(s) //[0 0 0]

}

func fillZero(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}

func fillZeroWithPointer(arr *[3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}

func fillZeroForSlice(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = 0
	}
}
