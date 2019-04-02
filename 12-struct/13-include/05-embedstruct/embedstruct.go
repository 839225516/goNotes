package main

import "fmt"

// 车轮
type Wheel struct {
	Size int
}

// 车
type Car struct {
	Wheel
	// 内嵌匿名结构体
	Engine struct {
		Power int
		Type  string
	}
}

func main() {

	c := Car{
		// 初始化车轮
		Wheel: Wheel{
			Size: 18,
		},

		//初始化引擎, 初始化时要先填写声明
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}

	fmt.Printf("%+v\n", c)
}
