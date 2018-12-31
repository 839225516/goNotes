package main

import "fmt"

func main() {

	var a = [3]int{1, 2, 3}

	fmt.Println(a, a[2:3])

	// 从数组或切片生产新的切片拥有如下特性
	// 取出的元素数量为：  结束位置 - 开始位置
	// 取出元素不包含结束位置对应的索引，切片最后一个元素使用slice[len(slice)]获取
	// 当缺省开始位置时，表示从开头到结束位置
	// 当缺省结束位置时，表示从开始位置到末尾
	// 两者都缺省时，与切片本身等效
	// 两者同时为0时，等效于空切片，一般用于切片复位
	// 根据索引位置取切片元素值时，取值范围是(0 ~ len(slice)-1 ),越界会报错
	// 生产切片时，结束位置可以填写 len(slice),但不会报错

	var highRiseBuilding [30]int

	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}

	//区间  highRiseBuilding[10:15]
	fmt.Println(highRiseBuilding[10:15])

	// 中间到尾部的所有元素 highRiseBuilding[20:]
	fmt.Println(highRiseBuilding[20:])

	// 开关到中间的所有元素 highRiseBuilding[:9]
	fmt.Println(highRiseBuilding[:9])

	// 表示原切片 highRiseBuilding[:]
	fmt.Println(highRiseBuilding[:])

	// 重置切片，清空拥有的元素
	var b = []int{1, 2, 3}
	fmt.Println(b[0:0])

}
