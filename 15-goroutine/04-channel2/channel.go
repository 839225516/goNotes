package main

import (
	"fmt"
	"time"
)

// 通道的数据接收可以借用for range语句进行多个元素的接收操作
// for data := range ch {}

func main() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		for i := 3; i >= 0; i-- {
			// 发送3 2 1 0
			ch <- i

			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

	// 遍历接收通道数据
	for data := range ch {
		fmt.Println(data)

		// 当遇到0 时，退出接收循环
		if data == 0 {
			break
		}
	}
}
