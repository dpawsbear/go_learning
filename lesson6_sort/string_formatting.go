package main

import (
	"fmt"
	"os"
)

type point struct {
	x,y int
}

func main(){
	p := point{1,2}
	fmt.Printf("%v\n" ,p)
	fmt.Printf("%+v\n",p)
	fmt.Printf("%#v\n",p)
	fmt.Printf("%T\n" ,p)

	fmt.Printf("%t\n",true)

	fmt.Printf("%d\n",123)

	//二进制形式打印
	fmt.Printf("%b\n",p)
	//字符串形式
	fmt.Printf("%c\n",33)
	//十六进制
	fmt.Printf("0x%x\n",456)
	//浮点
	fmt.Printf("%f\n",78.9)

	//不同的科学计数法
	fmt.Printf("%e\n",123400000.0)
	fmt.Printf("%E\n",123400000.0)

	//基本字符串输出
	fmt.Printf("%s\n","\"string\"")

	fmt.Printf("%q\n","\"string\"")
	//十六进制打印
	fmt.Printf("%x\n","hex this")
	//指针
	fmt.Printf("%p\n",&p)

	/*含间距格式化*/
	fmt.Printf("|%6d|%6d|\n",12,345)
	//使用 - 可以左对齐
	fmt.Printf("|%-6.2f|%-6.2f|\n",1.2,3.45)
	fmt.Printf("|%6.2f|%6.2f|\n",1.2,3.45)

	fmt.Printf("|%6s|%6s|\n","foo","b")
	fmt.Printf("|%-6s|%-6s|\n","foo","b")

	//sprintf也有的
	s := fmt.Sprintf("a %s","string")
	fmt.Println(s)

	//错误通道打印信息
	fmt.Fprintf(os.Stderr,"an %s\n","error")
}
