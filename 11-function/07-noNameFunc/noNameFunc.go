package main

import "fmt"

func main() {

	// 匿名函数定义格式， 就是没有名字的普通函数
	/*
		func(参数列表) (返回参数列表){
			函数体
		}
	*/

	// 匿名函数可以在声明后调用
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	// 将匿名函数赋值给变量
	f := func(data int) {
		fmt.Println("hello", data)
	}

	// 使用f()
	f(100)
}
