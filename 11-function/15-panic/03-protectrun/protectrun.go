package main

import (
	"fmt"
	"runtime"
)

// panic 和 defer 组合:
// 有panic 没 recover程序宕机
// 有panic 也有 recover ，程序不会宕机，执行完对应的 defer后，从宕机点退出当前函数后，继续执行

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string //所在函数
}

func ProtectRun(entry func()) {
	defer func() {
		// 宕机时，获取panic传递的上下文件并打印
		err := recover()

		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()

	entry()
}

func main() {
	fmt.Println("运行前")

	ProtectRun(func() {
		fmt.Println("手动触发panic前")

		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})

		fmt.Println("手动宕机后")
	})

	//故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")

		var a *int
		*a = 1

		fmt.Println("赋值宕机后")
	})

	fmt.Println("运行后")

}
