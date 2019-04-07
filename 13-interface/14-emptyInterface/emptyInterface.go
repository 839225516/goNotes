package main

import "fmt"

// 空接口类型---可以保存所有值的类型

func main() {
	var any interface{}

	// 空接口赋值
	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)

	var a int = 10

	var i interface{} = a

	// 这里会报错，不能将i视为int类型赋值给b
	//var b int = i

	// 使用类型断言
	var b int = i.(int)

	fmt.Println(b)

	// 空接口的值的比较
	var var1 interface{} = 100
	var var2 interface{} = "hi"

	// 两个空接口不相等
	fmt.Println(var1 == var2)
}
