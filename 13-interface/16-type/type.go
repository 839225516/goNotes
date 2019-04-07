package main

import "fmt"

/*
	//类型断言
	switch 接口变量.(type) {
	case 类型1:
		// 变量是类型1时的处理
	case 类型2:
		//变量是类型2时的处理
	...
	default:
		// 变量不是所有case列举中的类型时的处理
	}
*/

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, " is  int")
	case string:
		fmt.Println(v, " is string")
	case bool:
		fmt.Println(v, " is bool")
	default:
		fmt.Println(v, " is others")
	}
}

func main() {
	printType(1024)
	printType("hello go")
	printType(true)
}
