package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word = make(map[string]int64)
	inputs := bufio.NewScanner(os.Stdin)
	inputs.Split(bufio.ScanWords)
	for inputs.Scan() {
		word[inputs.Text()]++
	}
	if err := inputs.Err(); err != nil {
		fmt.Println(err)
	}
	for k, v := range word {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
