package main

import (
	"fmt"
	"time"
)

func main(){
	p := fmt.Println

	//得到当前时间
	nowtime := time.Now()
	p(nowtime)

	//创建一个时间结构体进行解析
	then := time.Date(
		2017,8,18,19,34,52,621544,time.UTC)

	p(then)

	//这里可以使用其他的函数不一一例举了
	p(then.Year())
	p(then.Month())
	p(then.Second())
	p(then.Location())
	p(then.Weekday())

	//同时可以进行时间的比对
	p(then.Before(nowtime))
	p(then.After(nowtime))
	p(then.Equal(nowtime))

	//两个时间的间隔
	diff := nowtime.Sub(then)

	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//时间加减
	p(then.Add(diff))
	p(then.Add(-diff))
}
