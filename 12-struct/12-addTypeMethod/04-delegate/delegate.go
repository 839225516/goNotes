package main

import "fmt"

type Class struct {
}

// 给结构体添加 Do() 方法
func (c *Class) Do(v int) {
	fmt.Println("call method do:", v)
}

// 普通函数的 Do() 方法
func funcDo(v int) {
	fmt.Println("call function do:", v)
}

func main() {
	// 声明一个函数回调
	var delegate func(int)

	// 创建结构体实例
	c := new(Class)

	// 将回调设为 c 的 Do 方法
	delegate = c.Do

	// 调用
	delegate(100)

	// 将回调设为普通函数
	delegate = funcDo

	// 调用
	delegate(100)
}
