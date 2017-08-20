package main

import (
	"os"
	"fmt"
)

func main(){
	//os.Args 提供了原始命令行参数的访问
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//可以单独提取参数
	arg := os.Args[3]

	fmt.Println(argsWithoutProg)
	fmt.Println(argsWithProg)
	fmt.Println(arg)
}
