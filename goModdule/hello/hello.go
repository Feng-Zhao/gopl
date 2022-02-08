package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// 设置log前缀
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// 初始化一个 slice 记录名字.
	names := []string{"Gladys", "Samantha", "Darrin"}
	// 调用 greetings 模块的 Hellos 方法
	messages, err := greetings.Hellos(names)
	// 检查异常/错误
	if err != nil {
		log.Fatal(err)
	}
	// 打印返回值
	fmt.Println(messages)

	//=============================================================
	// Hello 接受 一个 string 参数, 上边的的 Hellos 接受 []string 为参数

	// Request a greeting message.
	// message, err := greetings.Hello("Gladys")
	// // If an error was returned, print it to the console and
	// // exit the program.
	// if err != nil {
	// 	// log.Fatal(err) 内部使用 os.exit(1)退出程序
	// 	log.Fatal(err)
	// }

	// // If no error was returned, print the returned message
	// // to the console.
	// fmt.Println(message)
}
