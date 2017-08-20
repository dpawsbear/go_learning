package main

import (
	"flag"
	"fmt"
)

func main(){
	//基本标志声明可用于字符串，整数和布尔选项。这里我们声明一个带有默认值
	//foo 的字符串标识 word 和一个简短的描述。这个 flag.String 函数一个字符串指针
	wordPtr := flag.String("word","foo","a string")

	//声明 numb 和 fork标识，使用和word标识相同的方法
	numbPtr := flag.Int("numb",42,"an int")
	boolPtr := flag.Bool("fork",false,"a bool")

	//还可以声明一个使用在程序中其他地方声明的现有变量的选项。
	//我们需要传递一个指向标识声明的函数指针
	var svar string
	flag.StringVar(&svar,"svar","bar","a string var")

	//一旦所有的标识都声明好了，就可以解析了
	flag.Parse()


	//显示消息
	fmt.Println("word:",*wordPtr)
	fmt.Println("numb:",*numbPtr)
	fmt.Println("fork:",*boolPtr)
	fmt.Println("svar:",svar)
	fmt.Println("tail:",flag.Args())
}
