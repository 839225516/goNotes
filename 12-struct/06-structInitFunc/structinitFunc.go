package main

import "fmt"

// go 的类型或结构体没有构造函数的功能，初始化过程可以使用函数封装

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

func main() {
	var kity *Cat = NewCatByName("kitty")
	var whiteCat = &Cat{}
	whiteCat = NewCatByColor("white")

	fmt.Println(kity.Name)
	fmt.Println(whiteCat)
}
