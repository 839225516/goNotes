package main

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

type pig struct{}

func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	// 遍历映射
	for name, obj := range animals {

		// 判断是否为飞行动物, 对实例的interface{}变量进行断言操作
		f, isFlyer := obj.(Flyer)
		w, isWalker := obj.(Walker)

		fmt.Printf("name: %s; isFlyer: %v; is Walker: %v\n", name, isFlyer, isWalker)

		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}

}
