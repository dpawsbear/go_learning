package main

import (
	"time"
	"fmt"
)

func main(){

	//使用 time.Now() 与 Unix 或者 UnixNano 并用获得时间
	now  := time.Now()
	secs := now.Unix()
	nanos:= now.UnixNano()
	fmt.Println(now)


	millis := nanos/1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	//通过秒或者纳秒反向求解

	fmt.Println(time.Unix(secs,0))
	fmt.Println(time.Unix(0,nanos))
}
