package main

import (
	"os/exec"
	"os"
	"syscall"
)

func main(){
	//对于我们的例子，我们想要执行 notepad ,我们先用lookpath找到他
	binary,lookErr := exec.LookPath("notepad")
	if lookErr != nil {
		panic(lookErr)
	}

	//Exec 需要以 slice 形式的参数(作为一个大的字符串)
	args := []string{"notepad","a.c"}

	//exec 也需要一套环境来使用。这里我们只提供我们当前的环境。

	env := os.Environ()



	execErr := syscall.Exec(binary,args,env)
	if execErr != nil {
		panic(execErr)
	}
}