package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,
		"Path = %s \nMethod = %s \nContentLength = %d \nHost = %s \nProto = %s \nRemoteAddr = %s\n",
		r.URL.Path,
		r.Method,
		r.ContentLength,
		r.Host,
		r.Proto,
		r.RemoteAddr)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count = %d\n", count)
	mu.Unlock()
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		// 延迟 80 ms
		anim.Delay = append(anim.Delay, delay)
		// 加载画面
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
