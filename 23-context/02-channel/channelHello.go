package main

import "fmt"

func printHello(ch chan int) {
	fmt.Println("hello from printHello")
	ch <- 2
}

func main() {

	ch := make(chan int)

	go func() {
		fmt.Println("hello inline")
		ch <- 1
	}()

	go printHello(ch)

	fmt.Println("hello from main")

	i := <-ch

	fmt.Println("Recieved:", i)

	<-ch

}
