package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

//函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体
type FuncCaller func(interface{})

//实现Invoker的call
func (f FuncCaller) Call(p interface{}) {
	// 调用 f()函数体
	f(p)
}

func main() {
	var invoker Invoker

	// 将匿名函数转为 FuncCaller 类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	invoker.Call("hello")
}
