package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) Add(a Point) {
	p.X += a.X
	p.Y += a.Y
}

func (p *Point) Minus(a Point) {
	p.X -= a.X
	p.Y -= a.Y
}

func main() {
	a := Point{X: 1, Y: 1}
	b := Point{X: 3, Y: 3}
	// 方法表达式
	funcAdd := (*Point).Add
	// 第一个参数为接收者
	funcAdd(&a, b)
	fmt.Println(a)
}
