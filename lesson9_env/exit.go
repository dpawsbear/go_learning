package main

import (
	"fmt"
	"os"
)

func main(){
	defer fmt.Println("exit!")

	//退出状态码1
	os.Exit(1)
}
