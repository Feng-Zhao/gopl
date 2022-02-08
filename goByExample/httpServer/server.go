package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// 绑定处理方法 http.HandleFunc = func http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 启动 server
	http.ListenAndServe(":8090", nil)
}
