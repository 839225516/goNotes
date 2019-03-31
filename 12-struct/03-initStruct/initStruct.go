package main

import "fmt"

// 结构体可以使用 键值对 初始化字段
/*
	ins := 结构体类型名{
		字段1： 字段1的值,
		字段2:  字段2的值,
		...
	}
*/

type People struct {
	name  string
	child *People
}

func main() {
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	fmt.Println(relation.child.child.name)
}
