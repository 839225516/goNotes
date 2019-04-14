package main

import "fmt"

// 从已经关闭的通道接收数据或者正在接收数据时，将会接收到通道类型的零值，然后停止阻塞并返回

func main() {
	ch := make(chan int, 2)

	ch <- 0
	ch <- 1

	close(ch)

	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch

		fmt.Println(v, ok)
	}
}
