package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明一个退出用的通道
	exit := make(chan int)

	fmt.Println("start")

	// 1秒后调用匿名函数,时间到达后，匿名函数会在弄一个goroutine中被调用
	time.AfterFunc(time.Second, func() {
		// 1秒后，打印结果
		fmt.Println("one second after")

		// 通知main()的goroutine 又经有结果了
		exit <- 0
	})

	<-exit
}
