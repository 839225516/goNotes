package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// 函数支持多返回值
// 使用return语句返回时，值列表的顺序需要与函数声明的返回值类型一致
func tyepdTwoValues() (int, int) {
	return 1, 2
}

// 带有变量名的的返回值
// go支持对返回值进行命名，这样返回值就和参数一样拥有变量名和类型
// 命名的返回值为变量的默认值， 数值为 0；字符串为空字符串；bool 为 false; 指针为nil
func namedRetValues() (a, b int) {
	// 函数声明时将返回值命名为 a, b
	// 因此可以在函数体直接对函数返回值进行赋值
	// 在函数结束前要显示的使用 return 语句进行返回
	a = 1
	b = 2
	return
}

func named2RetValues() (a, b int) {
	a = 1
	return a, 3
}

func main() {
	// 函数声明，以 func 标识，后接着 函数名、参数列表、返回参数和函数体
	/*
		func 函数名(参数列表) (返回参数列表)｛
		｝
	*/

	fmt.Println(add(1, 3))

	fmt.Println(tyepdTwoValues())

	fmt.Println(namedRetValues())

	fmt.Println(named2RetValues())
}
