package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main(){
	//写入一个字符串到文件中
	d1 := []byte("hello \n go \n")
	err := ioutil.WriteFile("write_test.txt",d1,0644)
	check(err)

	//打开文件并写入数据（更为精细的调用）
	f ,err := os.Create("write_test2.txt")
	check(err)

	//打开文件后一般操作就是立即推迟关闭文件
	defer f.Close()

	//按照预期的方式写入字节片段
	d2 := []byte{115,111,109,101,10}
	n2,err := f.Write(d2)
	check(err)
	fmt.Printf("Write %d bytes \n",n2)

	//写入字符串
	n3 , err := f.WriteString("writes \n")
	fmt.Printf("write %d bytes \n",n3)

	//刷新缓冲，保证数据写到硬盘中去了
	f.Sync()

	//bufio 提高文件读写性能
	w := bufio.NewWriter(f)
	n4,err:=w.WriteString("buffered\n")
	fmt.Printf("write %d bytes\n",n4)

	w.Flush()
}
