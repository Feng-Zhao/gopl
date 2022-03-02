package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	slice := []int{5, 4, 3, 2, 1}
	Sort(slice)
	fmt.Println(slice)
}

func Sort(value []int) {
	var root *tree
	// 构造树
	for _, v := range value {
		root = add(root, v)
	}
	// 将树中元素按顺序加入 slice
	appendValues(value[:0], root)
}

// 将树中元素按顺序加入 slice
func appendValues(value []int, t *tree) []int {
	if t != nil {
		value = appendValues(value, t.left)
		value = append(value, t.value)
		value = appendValues(value, t.right)
	}
	return value
}

// 向树中添加元素
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
