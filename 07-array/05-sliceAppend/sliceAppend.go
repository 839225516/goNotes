package main

import "fmt"

func main() {
	// 声明一个整型切片
	var numbers []int

	for i := 0; i < 10; i++ {
		// 向 numbers 切片添加元素
		numbers = append(numbers, i)

		// len() 查看切片元素个数； cap() 查看切片的容量
		fmt.Printf("len: %d    cap: %d   pointer: %p \n", len(numbers), cap(numbers), numbers)
	}

	var car []string

	// append() 添加一个元素
	car = append(car, "OldDriver")

	//添加多个元素
	car = append(car, "Ice", "Sniper", "Monk")

	//添加切片
	team := []string{"Pig", "Flyingcake", "Chicken"}

	// team 后面加上 "..." ,表示将team整个添加到car的后面
	car = append(car, team...)

	fmt.Println(car)
}
