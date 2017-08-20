package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

// 读取文件需要检查大多数的调用错误
// 这个函数帮助简化错误检查

func check(e error){
	if e != nil {
		panic(e)
	}
}


func main(){
	//最基本的文件读取到内存中
	dat , err := ioutil.ReadFile("test.log")
	check(err)
	fmt.Print(string(dat))

	//通过得到文件句柄值，可以更好的控制文件的读取。
	f,err := os.Open("test.log")
	check(err)

	//读取到一个5字节的数组中
	b1 := make([]byte,5)
	n1,err:=f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n",n1,string(b1))

	//找到一个已知位置然后读取
	o2 , err := f.Seek(5,0)
	check(err)
	b2 := make([]byte,2)
	n2,err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2))

	//没有内置的倒带，但是 Seek(0,0) 可以实现这一点
	_,err = f.Seek(0,0)
	check(err)

	//bufio 包提供了一个缓冲的读取器，因为它具有许多小文件读取的效率，并且提供了额外的读取方式
	//适当使用可以提到程序性能。
	r4 := bufio.NewReader(f)
	b4,err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n",string(b4))

	//关闭文件
	f.Close()
}
