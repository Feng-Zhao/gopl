package main

import (
	"fmt"
	"os"
)

func main() {
	p1_1()
	p1_2()
	p1_3()
}

func p1_1() {
	fmt.Println("-------- 1.1 -----------")
	fmt.Println(os.Args[0])
}

func p1_2() {
	fmt.Println("-------- 1.2 -----------")
	for i, content := range os.Args[1:] {
		fmt.Printf("%d : %v \n", i, content)
	}
}
func p1_3() {
	fmt.Println("-------- 1.3 -----------")
	fmt.Println("not finished")
}
