package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个打点器，每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)

	// 创建一个计时器，2秒后触发
	stopper := time.NewTimer(time.Second * 2)

	// 声明计数变量
	var i int

	// 不断地检查通道情况
	for {
		// 多路复用通道
		select {
		case <-stopper.C:
			fmt.Println("stop")

			goto StopHere

		case <-ticker.C:
			// 记录触发了多少次
			i++
			fmt.Println("tick", i)

		}
	}

StopHere:
	fmt.Println("done")
}
