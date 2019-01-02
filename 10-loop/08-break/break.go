package main

import "fmt"

func main() {

	// break语句还可以在语句后面添加标签
	// 表示退出某个标签对应的代码块
	// 标签要求必须定义在对应的 for  switch 和 select 的代码块上
OutLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)

				//这里会退出 OutLoop 标记的循环
				break OutLoop
			case 3:
				fmt.Println(i, j)
				break OutLoop
			}

		}
	}
}
