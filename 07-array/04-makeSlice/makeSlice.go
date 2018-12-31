package main

import "fmt"

func main() {

	// make ( []T, size, cap)
	// size 分配多少个元素
	// cap 预分配的元素数量，这个值设定后不影响size,只是能提前分配空间

	// 使用 make()函数生成的切片一定发生了内存分配
	// 但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向了已经分配的内存区域
	// 设定开始与结束位置，不会发生内存分配

	// 切片不一定必须经过make()函数才能使用，生成切片、 声明后使用append()函数均可以正常使用切片

	// a和b都是预分配2个元素的切片，只是b的内部存储空间已经分配了10个，便实际使用了2个元素
	a := make([]int, 2)
	b := make([]int, 2, 10)

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
}
