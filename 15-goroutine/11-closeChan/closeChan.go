package main

import "fmt"

func main() {
	ch := make(chan int)

	// close chan
	// 关闭通道，不会把ch设置为nil,依然可以被访问
	close(ch)

	fmt.Printf("print: %p cap: %d; len: %d\n", ch, cap(ch), len(ch))

	// 给关闭的通道发送数据，会把panic
	// ch <- 1
}
