package main

import "fmt"

// 使用多值的列表初始化结构体
/*
	ins := 结构体类型名 {
		字段1的值,
		字段2的值,
		...
	}
*/

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

func main() {
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}

	fmt.Println(addr)
}
