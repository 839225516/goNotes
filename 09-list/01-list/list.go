package main

import (
	"container/list"
	"fmt"
)

func main(){
	// 列表没有具体元素类型的限制
	// list 的两种初始化方式
	// 通过container/list 包的 New 方法初始化
	List := list.New()

	// 通过声明初始化list
	var List2 list.List
	List2.PushBack("List2")

	// 双链表支持从列表前方或后方插入元素，对应的方法 PushFront 和 PushBack

	// 在列表尾部插入数据
	List.PushBack(67)

	// 在列表前插入数据，在 67前面插入 first
	List.PushFront("first")

	// 在67 插入 two,保存元素句柄
	element := List.PushBack("two")

	// 在two 之后添加 high
	List.InsertAfter("high",element)

	// 在 two 之前添加noon
	List.InsertBefore("noon",element)

	// 删除元素
	List.Remove(element)

	// 在List列表后插入List2
	List.PushBackList(&List2)

	// 遍历列表 遍历双链表要配合 Front() 函数取头元素，Nest()函数取下一元素
	for i:=List.Front(); i != nil; i= i.Next(){
		fmt.Println(i.Value)
	}
}
