package main

import "fmt"

type Walker interface {
	Walk()
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig： Walking")
}

func main() {
	//可以将实现的接口转换为普通类型的指针
	// 将Walker 转换为*pig
	p1 := new(pig)

	var a Walker = p1

	// 断言
	p2 := a.(*pig)

	fmt.Printf("p1=%p p2=%p", p1, p2)
}
