package main

import "fmt"

func main() {

	const elementCount = 1000

	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	// 引用切片数据
	// 切片不会因 = 进行元素的复制
	refData := srcData

	copyData := make([]int, elementCount)

	// 将数据复制到新的切片空间
	// copy(destSlice, srcSlice []T) int; 返回值表示实际发生复制的元素个数
	copy(copyData, srcData)

	//修改原始数据的第一个元素
	srcData[0] = 999

	//打印引用切片的第一个元素
	// 引用数据的第一个数据会改变
	fmt.Println(refData[0])

	// 打印复制切片的第一个和最后一个元素
	// 复制数据不会有变化
	fmt.Println(copyData[0], copyData[elementCount-1])

	// 复制原始数据从 4到6
	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}

}
