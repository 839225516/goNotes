package main

import (
	"fmt"
	"time"
)

// 使用匿名函数创建goroutine
/*
	go func(参数列表){
		函数体
	}(调用参数列表)
*/

func main() {
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)

			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
}
