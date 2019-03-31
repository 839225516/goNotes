package main

import "fmt"

// 使用type 定义新的自定义变量，之后就可以为自定义类型添加各种方法
type MyInt int

// 为MyInt添加Add()方法
func (m MyInt) IsZero() bool {
	return m == 0
}

func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero())

	b = 1
	fmt.Println(b.Add(2))
}
