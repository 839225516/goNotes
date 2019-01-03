package main

import "fmt"

func fire() {
	fmt.Println("fire")
}

func main() {

	// go 中函数也是一种类型，可以和其它类型一样保存在变量中
	// 声明变量f 为函数类型 func()
	var f func()

	f = fire

	f()
}
