package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	// recover() 必须在 defer func(){} 方法中
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	fmt.Println("Before mayPanic()")

	// 此处产生panic
	mayPanic()
	// 上方产生Panic，然后进入defer，这一句不会执行
	fmt.Println("After mayPanic()")
}
