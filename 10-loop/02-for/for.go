package main

import "fmt"

func main() {
	/*
		//for 循环格式：
		for 初始语句; 条件表达式; 结束语句 {
			循环体代码
		}
	*/

	// for 中的初始语句
	// 一般使用初始语句执行变量初初化，如果变量在此处被声明，其作用域被局限在这个for循环
	//初始语句可以被忽略，但初始语句后的分号 ";" 必须要写
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}

	// for 中的条件表达式，控制是否循环的开关
	// 每次循环开始前，计算表达式，如果表达式为 true，则循环继续，否则结束循环
	// 条件表达式可以被忽略，被忽略条件的表达式默认形成 无限循环
	var i int
	for ; ; i++ {
		fmt.Println(i)
		if i > 10 {
			break
		}
	}

	// 忽略for的所有语句，此时 for 为无限循环
	for {
		fmt.Println(i)
		if i > 13 {
			break
		}
		i++
	}

}
