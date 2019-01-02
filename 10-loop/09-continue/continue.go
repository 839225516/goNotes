package main

import "fmt"

func main() {
	// continue 可以结束当前循环，开始下一次的循环
	// 在continue 语句后添加标签，表示 开始标签对应的循环
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {

			switch j {
			case 2:
				fmt.Println(i, j)

				// 这里会开启外循环
				continue OuterLoop
			}
		}
	}
}
