package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i = 123456
	var err error
	if len(os.Args) == 2 {
		i, err = strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println(err)
			panic("error while convert number")
		}
	}
	fmt.Println(comma(strconv.Itoa(i)))
}

func comma(num string) string {
	if len(num) <= 3 {
		return num
	}
	return comma(num[:len(num)-3]) + "," + num[len(num)-3:]
}
