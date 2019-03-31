package main

import "fmt"

type Bag struct {
	items []int
}

// 将物品放入背包的过程
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

// 结构体方法
func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}

func main() {
	bag := new(Bag)

	Insert(bag, 100)

	bag.Insert(1001)

	fmt.Println(bag)
}
