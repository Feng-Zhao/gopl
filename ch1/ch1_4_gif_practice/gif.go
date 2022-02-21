package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// 练习 1.5, 添加更多颜色

// RGB 定义绿色
var green = color.RGBA{G: 0xff, A: 0xff}
var palette = []color.Color{color.Black, green}

const (
	blackIndex = 0
	greenIndex = 1
)

// go run gif.go web
// 然后访问 localhost:8000
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		// http 请求处理函数
		handler := func(w http.ResponseWriter, r *http.Request) {
			// 实际处理逻辑
			lissajous(w)
		}
		// http请求路径和处理函数的绑定
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // x 周期数
		res     = 0.001 // 角度分辨率
		size    = 100   // 画布大小 pix
		nframes = 64    //帧数
		delay   = 8     // 10ms 为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0        // y 振荡频率
	anim := gif.GIF{LoopCount: nframes} // gif 帧数
	phase := 0.0                        //相位
	// gif 动画 64 帧
	for i := 0; i < nframes; i++ {
		// 方形画布,大小 (0,0) 到 (201,201)
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 给画布配置色盘
		img := image.NewPaletted(rect, palette)
		// 绘制每帧动画
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 将 (x,y) 点设置为黑色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		// 延迟 80 ms
		anim.Delay = append(anim.Delay, delay)
		// 加载画面
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
