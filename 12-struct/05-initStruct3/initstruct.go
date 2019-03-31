package main

import "fmt"

// 匿名结构体： 没有类型名称，无须通过type 关键字定义就可以直接使用
// 初始化匿名结构体
/*
	ins := struct {
		// 匿名结构体字段定义
		字段1 字段类型1
		字段2 字段类型2
		...
	}{
		// 字段值初始化; 键值对初始化部分是可选的
		初始化字段1:  字段1的值
		初始化字段2： 字段2的值
		...
	}
*/

func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
}

func main() {
	// 实例化一个匿名结构体
	msg := &struct {
		// 定义部分
		id   int
		data string
	}{
		1024,
		"hello",
	}

	printMsgType(msg)
}
