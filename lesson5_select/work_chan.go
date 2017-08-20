package main

import (
	"fmt"
	"time"
)

//worker本体函数
func worker(id int,job <-chan int, result chan<- int){
	for j:=range job{
		fmt.Println("worker",id,"started job",j)
		time.Sleep(time.Second)
		fmt.Println("worker",id,"finished job",j)
		result<- j*2
	}
}

func main(){
	jobs:= make(chan int,100)
	results := make(chan int,100)

	//创建3个worker
	for w:=1 ; w<= 3;w++{
		go worker(w,jobs,results)
	}

	//分配5个任务
	for j:=1 ;j<= 5 ; j++{
		jobs <- j
	}

	close(jobs)

	//等待所有工作完成
	for a :=1 ; a<=5 ; a++{
		<- results
	}
}
