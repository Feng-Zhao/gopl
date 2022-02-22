package main

import (
	"fmt"
	"os"
	"strings"
	"time"
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

// 1.3 benchmark 在 echo_test 中
func p1_3() {
	fmt.Println("-------- 1.3 -----------")
	fmt.Println("not finished")

	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	start := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		fmt.Println(slice)
	}
	end := time.Now().Unix()
	diff := end - start

	start = time.Now().Unix()
	for i := 0; i < 100000; i++ {
		fmt.Println(strings.Join(slice, " "))
	}
	end = time.Now().Unix()
	fmt.Printf("1 time = %v秒", diff)
	fmt.Printf("2 time = %v秒", end-start)
}
