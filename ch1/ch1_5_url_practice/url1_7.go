package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// go run url.go http://www.baidu.com

func main() {
	// 从参数中获取 url
	for _, url := range os.Args[1:] {
		// 1.8
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		// Get 请求 url
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %c\n", err)
			os.Exit(1)
		}
		// 读取 body 部分 1.7
		reader := bufio.NewReader(resp.Body)
		if _, err = io.Copy(os.Stdout, reader); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s %v \n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
		// 1.9
		fmt.Printf("\n======== STATUD CODE ==========\n\t%v", resp.Status)
	}

}
