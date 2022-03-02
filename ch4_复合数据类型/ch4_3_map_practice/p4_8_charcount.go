package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type charType string

const digital charType = "digital"
const letter charType = "letter"
const other charType = "other"

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int // UTFMax 为 utf8 编码最大字节数
	charTypeCount := map[charType]int64{digital: 0, letter: 0, other: 0}
	invalid := 0

	// 从 stdin 读取
	in := bufio.NewReader(os.Stdin)
	for {
		// 读取一个 rune
		r, n, err := in.ReadRune()
		// 读取到 eof 结束读取
		if err == io.EOF {
			break
		}
		// 出错 code 1 退出
		if err != nil {
			fmt.Fprintf(os.Stdout, "charcount: %v\n", err)
			os.Exit(1)
		}
		// 读取到错误字符
		if r == unicode.ReplacementChar && n == 1 { //�
			invalid++
			continue
		}
		// 统计 rune, 和 每种字节长度 rune 的出现次数
		if unicode.IsDigit(r) {
			charTypeCount[digital]++
		} else if unicode.IsLetter(r) { // 这里中文字符也是算作 letter 的
			charTypeCount[letter]++
		} else {
			charTypeCount[other]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune \t count \n")
	for c, n := range counts {
		fmt.Printf("%q \t %d \n", c, n)
	}
	fmt.Printf("\nlen \t count \n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d \t %d \t \n", i, n)
		}
	}
	fmt.Printf("\ntype\tcount\n")
	for k, v := range charTypeCount {
		fmt.Printf("%s\t%d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
