package main

import "fmt"

func main() {

	// 闭包是引用了外部变量的匿名函数

	str := "hello World"

	foo := func() {

		// 匿名函数中访问 str
		str = "hello Dude"
		fmt.Println(str)
	}

	foo()
	fmt.Println("str:", str)
}
