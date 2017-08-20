package main

import "os"

func main(){

	panic("a problem")

	_,err := os.Create("./test.log")
	if err != nil{
		panic(err)
	}
}