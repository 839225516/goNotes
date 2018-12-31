package main

import (
	"fmt"
	"sort"
)

func main() {
	sence := make(map[string]int)

	sence["route"] = 66
	sence["ZZ"] = 96
	sence["CZ"] = 98

	// 声明一个切片保存 map 数据
	var senceList []string

	// 将map数据遍历复制到切片中
	// 将map中元素的键遍历，并放入切片中
	for k := range sence {
		senceList = append(senceList, k)
	}

	// 对切片进行排序，排序时 senceList 会被修改
	sort.Strings(senceList)

	// 输出排好序的键
	fmt.Println(senceList)
}
