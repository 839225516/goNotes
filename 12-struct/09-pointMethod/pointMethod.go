package main

import "fmt"

type Property struct {
	value int
}

func (p *Property) SetValue(v int) {
	p.value = v
}

func (p *Property) Value() int {
	return p.value
}

func main() {
	// 实例化属性
	p := new(Property)

	p.SetValue(1000)

	fmt.Println(p.Value())
}
