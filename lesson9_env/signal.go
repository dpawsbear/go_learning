package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	//Go 信号通知通过在通道上发送 os.Signal 值
	//我们将创建一个频道来接收这些通知
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
