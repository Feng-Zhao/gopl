package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// 用select模拟server长时间没有反馈
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	// 当客户端取消访问时，结束请求
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
