package main

import (
	"container/list"
	"fmt"
)

func main(){

	// func (e *Element) Next() *Element  //返回该元素的下一个元素，如果没有下一个元素则返回nil
	// func (e *Element) Prev() *Element//返回该元素的前一个元素，如果没有前一个元素则返回nil
	// func New() *List //返回一个初始化的list (双链表)
	// func (l *List) Back() *Element //获取list l的最后一个元素
	// func (l *List) Front() *Element //获取list l的第一个元素
	// func (l *List) Init() *List  //list l初始化或者清除list l
	// func (l *List) InsertAfter(v interface{}, mark *Element) *Element  //在list l中元素mark之后插入一个值为v的元素，并返回该元素，如果mark不是list中元素，则list不改变
	// func (l *List) InsertBefore(v interface{}, mark *Element) *Element//在list l中元素mark之前插入一个值为v的元素，并返回该元素，如果mark不是list中元素，则list不改变
	// func (l *List) Len() int //获取list l的长度
	// func (l *List) MoveAfter(e, mark *Element)  //将元素e移动到元素mark之后，如果元素e或者mark不属于list l，或者e==mark，则list l不改变
	// func (l *List) MoveBefore(e, mark *Element)//将元素e移动到元素mark之前，如果元素e或者mark不属于list l，或者e==mark，则list l不改变
	// func (l *List) MoveToBack(e *Element)//将元素e移动到list l的末尾，如果e不属于list l，则list不改变
	// func (l *List) MoveToFront(e *Element)//将元素e移动到list l的首部，如果e不属于list l，则list不改变
	// func (l *List) PushBack(v interface{}) *Element//在list l的末尾插入值为v的元素，并返回该元素
	// func (l *List) PushBackList(other *List)//在list l的尾部插入另外一个list，其中l和other可以相等
	// func (l *List) PushFront(v interface{}) *Element//在list l的首部插入值为v的元素，并返回该元素
	// func (l *List) PushFrontList(other *List)//在list l的首部插入另外一个list，其中l和other可以相等
	// func (l *List) Remove(e *Element) interface{}//如果元素e属于list l，将其从list中删除，并返回元素e的值

	List := list.New()	// 创建一个新的list
	for i:=0; i<5 ; i++{
		List.PushBack(i)	// 在List尾部依次插入元素
	}

	for e:=List.Front(); e!=nil; e=e.Next(){
		fmt.Println(e.Value)		//输出List的值
	}

	// List首部元素的值
	fmt.Println("首部元素的值是：",List.Front().Value)

	// 尾部元素的值
	fmt.Println("尾部元素的值是：",List.Back().Value)

	// 在首部元素之后插入一个值6
	List.InsertAfter(6, List.Front())
	for e:=List.Front(); e!=nil; e=e.Next(){
		fmt.Println(e.Value)		//输出List的值
	}

	// 将首部的两个元素互换
	List.MoveBefore(List.Front().Next(), List.Front())
	for e:=List.Front(); e!=nil; e=e.Next(){
		fmt.Print(e.Value)		//输出List的值
	}
	fmt.Println("")
	// 将尾部元素移动到首部
	List.MoveToFront(List.Back())
	for e:=List.Front(); e!=nil; e=e.Next(){
		fmt.Println(e.Value)		//输出List的值
	}

	// List的长度
	fmt.Println("List的长度是：",List.Len())
}
