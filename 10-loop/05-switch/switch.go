package main

import "fmt"

func main() {
	// switch 每个case与case间是独立的代码块，不需要通过break语句跳出当前case
	// 每个switch 只能有一个 default
	var a string = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	// 一分支多值
	// 当出现多个case要放在一起时， 不同的表达式用逗号分隔
	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}

	// 分支表达式
	// case 后不仅仅只是常量，还可以是表达式
	// 这种情况 switch 后面不用再跟判断变量
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	// 跨越case 的 fallthrough ; 兼容 C 语言的case设计
	// go中 case 是一个独立的代码块 执行完毕后不会像C那样紧接着下一个case执行。
	// 但为了兼容移植代码，依然加入了 fallthrough 关键字 来实现这一功能
	var s string = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}

}
