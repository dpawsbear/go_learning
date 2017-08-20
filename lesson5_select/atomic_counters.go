package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

func main(){
	var ops uint64 = 0


	for i:= 0 ; i< 50 ; i++{
		go func() {
			for{
				atomic.AddUint64(&ops,1)
				time.Sleep(time.Millisecond * 1 )
			}
		}()
	}

	time.Sleep(time.Second * 10)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops",opsFinal)

}