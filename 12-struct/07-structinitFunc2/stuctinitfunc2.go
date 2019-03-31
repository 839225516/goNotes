package main

import "fmt"

type Cat struct {
	Color string
	Name  string
}

type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

// 构造基类
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 构造子类
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

func main() {
	var kity *Cat = NewCat("kitty")
	var whiteCat = &BlackCat{}
	whiteCat = NewBlackCat("black")

	fmt.Println(kity.Name)
	fmt.Println(whiteCat)
}
