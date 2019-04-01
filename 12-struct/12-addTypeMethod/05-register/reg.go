package main

import "fmt"

// 实例化一个通过字符串映射函数切片的 map
var eventByName = make(map[string][]func(interface{}))

// 注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	// 通过名字查找事件列表
	list := eventByName[name]

	// 在列表切片中添加函数
	list = append(list, callback)

	// 保存修改的事件列表切片
	eventByName[name] = list
}

// 事件调用
func CallEvent(name string, param interface{}) {
	// 通过名字找到事件列表
	list := eventByName[name]

	// 遍历这个事件的所有回调
	for _, callback := range list {
		// 传入参数调用回调
		callback(param)
	}
}

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

// 全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {
	// 实例化一个角色
	a := new(Actor)

	// 注册名为 OnSkill 的回调
	RegisterEvent("OnSkill", a.OnEvent)

	// 再次在OnSkill上注册全局事件
	RegisterEvent("OnSkill", GlobalEvent)

	// 调用事件，所有注册的同名函数都会被调用
	CallEvent("OnSkill", 100)
}
