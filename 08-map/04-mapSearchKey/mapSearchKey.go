package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["中国"] = "北京"
	countryCapitalMap["India"] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country,"的首都是：",countryCapitalMap[country])
	}

	// 查看元素在 map的 key 中是否存在
	// 查看 美国 是否在 countryCapitalMap的 key 中
	capital, ok := countryCapitalMap["美国"]
	//fmt.Println(capital,ok)
	if ok {
		fmt.Println("美国的首都是", capital)
	} else {
		fmt.Println("没找到美国的首都")
	}
}
