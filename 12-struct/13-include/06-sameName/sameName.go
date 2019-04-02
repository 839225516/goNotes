package main

import "fmt"

type A struct {
	a int
}

type B struct {
	// 与A 成员同名
	a int
}

type C struct {
	A
	B
}

func main() {
	c := &C{}

	// c.a = 1 //这样会报错

	c.A.a = 1
	fmt.Println(c)
}
