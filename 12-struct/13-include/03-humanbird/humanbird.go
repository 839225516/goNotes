package main

import "fmt"

// 可飞行的
type Flying struct {
}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// 可行走的
type Walkable struct{}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

// 人类
type Human struct {
	Walkable
}

type Bird struct {
	Walkable
	Flying
}

func main() {
	// 实例化鸟类
	b := new(Bird)
	fmt.Println("Bird: ")
	b.Fly()
	b.Walk()

	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()

}
