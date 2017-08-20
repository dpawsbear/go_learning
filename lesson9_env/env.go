package main

import (
	"os"
	"fmt"
	"strings"
)

func main(){
	//要设置 key/value 对，使用 os.Setenv。
	// os.Getenv 可以获取一个键值，如果没有会返回空


	os.Setenv("FOO","1")
	fmt.Println("Foo:",os.Getenv("FOO"))
	fmt.Println("BAR:",os.Getenv("BAR"))

	//使用 `os.Environ` 列出环境中所有的key/value对
	//可以使用 "string.Split"来获取键值和值，这里我们列举所有key
	fmt.Println()
	for _,e:=range os.Environ(){
		pair := strings.Split(e,"e")
		fmt.Println(pair[0])
	}
}
