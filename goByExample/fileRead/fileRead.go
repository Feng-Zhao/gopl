package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 读文件到 []byte
	dat, err := os.ReadFile("./dat.txt")
	check(err)
	fmt.Println(string(dat))

	// 打开文件 得到文件指针 *os.File
	f, err := os.Open("./dat.txt")
	// 关闭 文件
	check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	// 读取到缓存区 []byte
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes:\n%s\n", n1, string(b1[:n1]))

	// 移动 offset
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	// 使用 bufio.Reader 读取
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}
