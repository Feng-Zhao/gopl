package main

import "fmt"

func main() {
	fmt.Println("==== range slice ====")
	rangeSlice()
	fmt.Println("==== range map ====")
	rangeMap()
	fmt.Println("==== range channel ====")
	channelRange()
}

// 对 slice/array 使用 range 遍历， 有顺序
func rangeSlice() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
}

// 对 map 使用 range 遍历 注：for range map 的顺序是随机的
func rangeMap() {
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range for string, 遍历的是 unicode 的 code point
	fmt.Println("==== range string 'go' ====")
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

// range for channel 保持顺序
func channelRange() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
