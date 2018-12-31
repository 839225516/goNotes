package main

import "fmt"

func main() {
	// 声明数组,使用时可以修改成员，但数组大小不可以变化
	// var 数组变量名   [元素数量]T
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"

	fmt.Println(team)

	//初始化数组
	// 需要保证大括号后面的元素与数组的大小一致
	var array = [3]string{"hammer", "soldier", "mub"}
	fmt.Println(array)

	// ... 表示让编译器确定数组大小
	var array2 = [...]string{"hammer", "soldier", "mum"}
	fmt.Println(array2)

	// 遍历数组
	for k, v := range array {
		fmt.Println(k, v)
	}

}
