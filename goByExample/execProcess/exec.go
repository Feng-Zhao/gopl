package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 查找 ls 命令的路径 "/bin/ls"
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()
	// 这里执行 ls 命令
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	// 无法到达的代码，上边 Exec 会把程序控制权交给 ls 命令， 这一点就是
	// syscall.Exec(binary, args, env) 和 exec.Command("ls") 不同的地方
	fmt.Println("cannot get here")
}
