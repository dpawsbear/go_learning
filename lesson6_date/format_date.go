package main

import (
	"fmt"
	"time"
)

func main(){
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))


	t1 , _ := time.Parse(
		time.RFC3339,
		"2017-11-01T22:35:52+00:00")
	p(t1)

	//格式化和解析

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 20017"))
	p(t.Format("2006-01-02T15:04:05.99999-08:00"))

	form := "3 04 PM"
	t2,_ := time.Parse(form,"8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
			t.Year(),t.Month(),t.Day(),
			t.Hour(),t.Minute(),t.Second())

	//错误处理
	ansic := "Mon jan _2 15:04:05 2006"
	_,e:=time.Parse(ansic,"8:41PM")
	p(e)
}
