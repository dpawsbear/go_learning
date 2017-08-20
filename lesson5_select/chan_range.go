package main

import "fmt"

func main(){
	//存放两个值在通道中

	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem:=range queue{
		fmt.Println(elem)
	}
}
