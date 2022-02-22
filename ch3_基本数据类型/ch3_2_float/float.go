package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6 //30°
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill:white;stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			h := int(dy*sin30) * 256
			r, g, b := 0, 0, 0
			if h > 0 {
				r = h
			} else {
				b = -h
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='rgb(%d,%d,%d)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Println("</svg>")
}
func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}

// 3.2 变更 f 函数, 即,不同的 z = f(x,y)
func f1(x, y float64) float64 {
	z := math.Pow(x, 2) / (math.Pow(y, 2) + 0.1)
	return z
}
