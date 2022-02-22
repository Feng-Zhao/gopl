package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

// 初始化 8 位数中的 1 的数量
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "usage go run popcount.go n, n is an integer")
		os.Exit(1)
	}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "arg is not an integer")
		os.Exit(1)
	}
	fmt.Printf("%d\n", Popcount(uint64(x)))
}

func Popcount(x uint64) int {
	// 每次统计 8 位
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
