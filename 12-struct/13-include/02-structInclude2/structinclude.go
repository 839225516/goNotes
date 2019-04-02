package main

import "fmt"

type BasicColor struct {
	R, G, B float32
}

// 使用内嵌方法
type Color struct {
	// BasicColor 没有字段名，只有类型，这种写法叫结构体内嵌
	BasicColor
	Alpha float32
}

func main() {
	var c Color
	c.R = 1
	c.G = 1
	c.B = 0

	c.Alpha = 1

	fmt.Printf("%+v", c)
}
