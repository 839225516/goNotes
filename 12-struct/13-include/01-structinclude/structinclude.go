package main

import "fmt"

type BasicColor struct {
	R, G, B float32
}

// 完整颜色定义
type Color struct {
	// 将基本颜色定义为成员
	Basic BasicColor

	// 透明度
	Alpha float32
}

func main() {
	var c Color

	// 设置基本颜色分量
	c.Basic.R = 1
	c.Basic.G = 1
	c.Basic.B = 0

	// 设置透明度
	c.Alpha = 1

	fmt.Printf("%+v", c)
}
