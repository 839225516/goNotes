package main

import "fmt"

func main() {
	//定义map:  map[Keytype]ValueType
	// map是一个内部实现的类型，使用时，需要手动使用make创建
	// 如果不用make创建，使用时会触发宕机panic错误
	sence := make(map[string]int)

	// 向map中加入映射关系
	sence["route"] = 66

	fmt.Println(sence["route"])

	// 尝试查找一个不存在的键，那么返回的将是ValueType的默认值
	v := sence["route2"]
	fmt.Println(v)

	sence["brazi"] = 4
	sence["zz"] = 96

	// 查看 map的容量
	fmt.Println("sence的容量是",len(sence))

	// 遍历map, for range
	for k, v := range sence {
		fmt.Println(k, v)
	}

	// 只遍历值
	for _, value := range sence {
		fmt.Println(value)
	}

	// 只遍历键
	for k := range sence {
		fmt.Println(k)
	}
}
