package main

import (
	"fmt"
	"sort"
)

/*
	go 1.8 开始，在sort包提供了 sort.Slice()
	sort.Slice()函数只要传入需要排序的数据，以及一个排序时对元素的回调函数
	回调函数类型为 func(i,j int)

	sort.Slice()的定义如下：
	func Slice(slice interface{}, less(i,j int) bool)
*/

type HeroKind int

const (
	None = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	kind HeroKind
}

func main() {
	heros := []*Hero{
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲已", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].kind != heros[j].kind {
			return heros[i].kind < heros[j].kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
