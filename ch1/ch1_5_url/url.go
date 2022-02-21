package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// go run url.go http://www.baidu.com

func main() {
	// 从参数中获取 url
	for _, url := range os.Args[1:] {
		// Get 请求 url
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %c\n", err)
			os.Exit(1)
		}
		// 读取 body 部分
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s %v \n", url, err)
			os.Exit(1)
		}
		// 打印
		fmt.Printf("%s\n", b)
	}

}
