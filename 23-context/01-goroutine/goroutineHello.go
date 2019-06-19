package main

import "fmt"

func printHello() {
	fmt.Println("hello world from printHello")
}

func main() {
	// goroutine
	go func() { fmt.Println("hello inline") }()

	// goroutine
	go printHello()

	// 运行后只会打印这句
	fmt.Println("hello from main")
}
