package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 创建 临时文件 “” 代表操作系统默认的存放临时文件的位置
	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())
	// 删除临时文件
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 创建临时文件夹
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)
	// 删除临时文件夹
	defer os.RemoveAll(dname)

	// 向临时文件夹内的文件进行写入
	fname := filepath.Join(dname, "file1")
	// 文件不存在时会自动创建文件
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
