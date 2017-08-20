package main

import (
	"strconv"
	"fmt"
)

func main(){

	//后面的64是 64bits 的意思
	f,_ := strconv.ParseFloat("1.456",64)
	fmt.Printf("%T %f\n",f,f)

	//0 表示使用进制解析，修改为8或者16就按照8进制，16进制解析
	// 64表示结果必须符合 64bit（只能在INT64范围内）
	i,_ := strconv.ParseInt("123",0,64)
	fmt.Printf("%T %d\n",i,i)

	//ParseInt 可以识别十六进制的数字
	d,_:=strconv.ParseInt("0x1c8",0,64)
	fmt.Printf("%T %d\n",d,d)

	u,_:=strconv.ParseUint("789",0,64)
	fmt.Printf("%T %d\n",u,u)

	//atoi 基本的十进制解析
	k,_:=strconv.Atoi("135")
	fmt.Printf("%T %d\n",k,k)
	//错误
	_,e:= strconv.Atoi("test")
	fmt.Println(e)
}
