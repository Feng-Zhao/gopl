package main

import "fmt"

const UNIT_CAP = 64

func main() {
	set := &IntSet{words: make([]uint64, 32, 128)}
	set.Add(1)
	set.Add(149)
	fmt.Println(set)
	fmt.Println(set.Has(3))
	fmt.Println(*set)
	set.Add(0)
	fmt.Println(set)
	// String() 方法是绑定 *IntSet 的, IntSet 没有此方法,会按照 []uint64 形式打印
	fmt.Println(*set)
}

/*
	{1 149}
	false
	{[2 0 2097152 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]}
	{0 1 149}
	{[3 0 2097152 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]}
*/
