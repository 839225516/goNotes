package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
	}

	// 结构体成员变量的赋值
	var p Point
	p.X = 10
	p.Y = 20

	fmt.Printf("p.X = %d, p.Y = %d\n", p.X, p.Y)
}
