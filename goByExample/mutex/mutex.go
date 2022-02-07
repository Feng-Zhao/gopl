package main

import (
	"fmt"
	"sync"
)

type Container struct {
	// 使用 mutex 控制同步操作
	mu       sync.Mutex
	counters map[string]int
}

// 注意：传递 mutex 时应使用指针
func (c *Container) inc(name string) {
	// 加锁
	c.mu.Lock()
	// 操作结束之后解锁
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// 定义函数变量
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// wait group
	wg.Add(3)
	// start 3 goroutine
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
