// 2.4 + 2.5 练习
package ch2_6_popcount_practice

import (
	"math/rand"
	"testing"
)

var pc [256]byte

// 初始化 8 位数中的 1 的数量
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//func main() {
//	if len(os.Args) != 2 {
//		fmt.Fprint(os.Stderr, "usage go run popcount.go n, n is an integer")
//		os.Exit(1)
//	}
//	x, err := strconv.Atoi(os.Args[1])
//	if err != nil {
//		fmt.Fprint(os.Stderr, "arg is not an integer")
//		os.Exit(1)
//	}
//	fmt.Printf("%d\n", Popcount(uint64(x)))
//}

// 书中给出的 8 位一统计的方式
func Popcount1(x uint64) int {
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

// for 64 统计每位的方式
func Popcount2(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x = x >> 1
	}
	return count
}

// x&(x-1)清除末尾位,
func Popcount3(x uint64) int {
	count := 0
	for ; x != 0; x = x & (x - 1) {
		count++
	}
	return count
}

func Benchmark1(b *testing.B) {
	rand.Seed(1994)
	for i := 0; i < b.N; i++ {
		Popcount1(rand.Uint64())
	}
}

func Benchmark2(b *testing.B) {
	rand.Seed(1994)
	for i := 0; i < b.N; i++ {
		Popcount2(rand.Uint64())
	}
}

func Benchmark3(b *testing.B) {
	rand.Seed(1994)
	for i := 0; i < b.N; i++ {
		Popcount3(rand.Uint64())
	}
}
