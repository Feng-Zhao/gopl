package main

import "os"

func main() {

	// panic 之后程序会异常退出
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
