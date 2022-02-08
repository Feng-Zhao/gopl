package main

import (
	"fmt"
	"time"
)

// 返回 Unix 时间
func main() {

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// 第一个参数为 秒， 第二个参数为 纳秒
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

// 2022-02-08 15:55:56.433717 +0800 CST m=+0.000089516
// 1644306956
// 1644306956433
// 1644306956433717000
// 2022-02-08 15:55:56 +0800 CST
// 2022-02-08 15:55:56.433717 +0800 CST
