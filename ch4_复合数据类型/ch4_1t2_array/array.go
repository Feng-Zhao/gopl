package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var a = [3]int{1, 2, 3}
	modifyArrayBackup(a)
	fmt.Println("数组的参数传递为值传递,会复制一个副本供函数内部使用")
	fmt.Println(a)
	modifyArrayPointer(&a)
	fmt.Println("显式使用指针传递,函数内部对数组的修改会反馈在外部")
	fmt.Println(a)
	fmt.Println("==========练习 4.1========")
	sha41([]byte("francis"))
}
func modifyArrayBackup(a [3]int) {
	a[0] = 0
	a[1] = 0
	a[2] = 0
}
func modifyArrayPointer(a *[3]int) {
	a[0] = 0
	a[1] = 0
	a[2] = 0
}

// 练习 4.1 统计 sha256 摘要中 1 的位数
// 4.2 针对输入输出 sha256 摘要
func sha41(b []byte) int {
	fmt.Printf("输入为: %s", b)
	fmt.Println("byte 数组形式: ", b)
	c := sha256.Sum256(b)
	fmt.Println("sha256 摘要为: ", c)
	fmt.Printf("16 进制: %x\n", c)
	num := shaCount(c)
	fmt.Println("一共", num, "个 1")
	return num
}
func shaCount(b [32]byte) int {
	count := 0
	for _, v := range b {
		count += bitCount(v)
	}
	return count
}
func bitCount(b byte) int {
	count := 0
	for ; b != 0; b = b & (b - 1) {
		count++
	}
	return count
}
