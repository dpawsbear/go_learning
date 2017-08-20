package main

import (
	"fmt"
	"time"
)

func main(){
	messages := make(chan string)
	signals  := make(chan bool)

	go func() {
		messages <- "test"
	}()

	time.Sleep(time.Second * 1)
	select{
		case msg := <-messages:
			fmt.Println("Received message",msg)
		default:
			fmt.Println("no message received")
	}

	msg := "hi"

	select {
		case messages<-msg:
			fmt.Println("sent message",msg)
		default:
			fmt.Println("no message sent")
	}

	select {
		case msg := <- messages:
			fmt.Println("received message",msg)
		case sig := <-signals:
			fmt.Println("received signals",sig)
		default:
			fmt.Println("no activity")
	}
}
