package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// 通过exec调用其他程序/命令
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 创建一个 grep
	grepCmd := exec.Command("grep", "hello")
	// 设置输入输出流
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	// 启动命令
	grepCmd.Start()
	// 写入内容
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	// 结束写入
	grepIn.Close()
	// 获取输出
	grepBytes, _ := io.ReadAll(grepOut)
	// 等待命令结束
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 使用 bash -c + 命令字符串直接调用
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
