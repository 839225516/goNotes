package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	// go 函数调用格式：
	// 返回值变量列表 = 函数名（参数列表）
	result := add(1, 1)
	fmt.Println(result)

}
