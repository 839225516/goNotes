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

}
