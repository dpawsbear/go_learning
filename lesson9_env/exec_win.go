package main

import (
"os/exec"
"fmt"
)

func main(){
	c := exec.Command("cmd","/C","notepad","a.c")
	if err := c.Run(); err != nil{
		fmt.Println("Error:",err)
	}
}
