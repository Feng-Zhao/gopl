package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 准备写入的内容
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("./dat1.txt", d1, 0644)
	check(err)

	f, err := os.Create("./dat2.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	// 写 []byte
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// 写 string
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	// 把缓冲区 flush 到存储介质
	f.Sync()

	// 使用 bufio.Writer 写入
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	// bufio.Writer flush 到存储介质
	w.Flush()

}
