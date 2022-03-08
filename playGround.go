package main

import (
	"fmt"
)

type S struct {
	X int
}

func main() {
	// play1()
}

func play1() {
	x := ^uint(0)
	fmt.Printf("%#b\n", x)

	slice := []int{0, 0, 0, 0}
	fmt.Println(slice) // [0 0 0 0]
	for i := range slice {
		slice[i] = 1
	}
	fmt.Println(slice) // 通过索引修改基本类型,有效果 [1 1 1 1]

	for _, v := range slice {
		v += 1
	}
	fmt.Println(slice) // 通过值修改 无效果 [1 1 1 1]

	slice2 := make([]S, 3, 3)
	fmt.Println(slice2) // [{0} {0} {0}]
	slice2[0] = S{X: 1}
	slice2[1] = S{X: 1}
	slice2[2] = S{X: 1}
	fmt.Println(slice2) // 直接修改,有效果 [{1} {1} {1}]

	for i := range slice2 {
		slice2[i].X = 2
	}
	fmt.Println(slice2) // 通过索引修改,有效果 [{2} {2} {2}]

	for _, v := range slice2 {
		v.X = 3
	}
	fmt.Println(slice2) // 直接修改值,无效果 [{2} {2} {2}]

}
