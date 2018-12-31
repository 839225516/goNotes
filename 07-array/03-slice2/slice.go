package main

import "fmt"

func main() {

	// 声明切片
	// var name []T
	// 声明字符串切片
	var strList []string

	// 声明整形切片
	var numList []int

	// 声明一个空切片
	var numListEmpty = []int{}

	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)

	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))

	// 判断切片是否为空
	// 声明了但未使用的切片的默认值是 nil
	// numListEmpty 已经被分配到了内存，但没有元素，因此不为 nil
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}
