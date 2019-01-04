package main

import "fmt"

// 调用器接口
type Invoke interface {
	// 需要实现一个Call()方法
	Call(interface{})
}

// 结构体实现接口

// 结构体类型
type Struct struct {
}

// 实现Invoke 的 Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("form struct ", p)
}

func main() {
	// 接口声明
	var invoke Invoke

	// 实例化结构体
	s := new(Struct)

	// 将实体化的结构体赋值到接口
	invoke = s

	invoke.Call("hello")

}
