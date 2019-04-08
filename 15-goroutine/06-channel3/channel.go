package main

import (
	"fmt"
	"time"
)

/*
	单向通道
		chan<- 只能发送的通道
		<-chan 只能接收的通道
*/

func Echo(out chan<- string) {
	time.Sleep(time.Second)
	out <- "hello goroutine"
	close(out)
}

func Receive(out chan<- string, in <-chan string) {
	temp := <-in
	out <- temp
	close(out)
}

func main() {
	echo := make(chan string)
	receive := make(chan string)

	go Echo(echo)
	go Receive(receive, echo)

	getStr := <-receive
	fmt.Println(getStr)
}
