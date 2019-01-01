package main

import (
	"fmt"
)

func main() {
	// for rang 可以遍历 数据、切片、字符串、map及通道channel
	// 数组、切片、字符串返回索引和值
	// map 返回键和值
	// 通道channel 只返回通道内的值

	// 数组、切片
	for key, value := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("key: %d   value: %d\n", key, value)
	}

	// 字符串
	var str = "hello ,你好啊！"
	for key, value := range str {
		fmt.Printf("key: %d  value: %c\n", key, value)
	}

	// map
	m := map[string]int{
		"hello": 10,
		"world": 200,
	}

	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	// channel 接收通道数据
	c := make(chan int)

	// 向通道中推送数据 1，2，3； 然后结束并关闭通道
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	// 从通道中取数据，直到通道被关闭
	for v := range c {
		fmt.Println(v)
	}

}
