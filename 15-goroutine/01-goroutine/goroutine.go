package main

import (
	"fmt"
	"time"
)

// 为普通函数创建goroutine的写法
// go 函数名(参数列表)

func running() {

	var times int

	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)

		// 延时一秒
		time.Sleep(time.Second)
	}
}

func main() {
	// 并发执行程序
	go running()

	// 接受命令行输入，不做任何事情
	// 等待用户输入，按回车后结束输入的内容，同时整个程序也终于了
	var input string
	fmt.Scanln(&input)
}
