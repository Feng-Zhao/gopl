package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -1, -1, +1, +1
		width, height          = 2048, 2048
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	//img := image.NewYCbCr(image.Rect(0, 0, width, height), 1)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iteration = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iteration; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			// 3.5 有色彩的图像
			return color.YCbCr{Y: n, Cb: 255 - contrast*n, Cr: 255 - contrast*n}
		}
	}
	return color.Black
}
