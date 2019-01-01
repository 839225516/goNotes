package main

import (
	"fmt"
	"sync"
)

func main(){

	// sync.Map 是sync包下的特殊结构
	// sync 无须初始化，直接声明可用
	// sync.Map 不能使用map的方式进行取值和设置等操作
	// 使用sync.Map的方法进行调用： Store 表示存储； Load 表示获取；Delete表示删除
	// 使用Range配置一个回调函数进行遍历操作
	// Range 参数中的回调函数的返回值功能是：需要继续迭代遍历时，返回 true；终止迭代遍历时，返回 false
	// sync.Map 没有提供获取 map 数量的方法

	var scene sync.Map

	//将键值对保存到sync.Map
	scene.Store("greece",97)
	scene.Store("london",100)
	scene.Store("egypt",200)

	//从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))

	// 根据键删除对应的键值
	scene.Delete("london")

	// 遍历所有的 sync.Map 中的键值
	scene.Range(func(k,v interface{}) bool{
		fmt.Println("iterate:", k, v)
		return true
	})
}
