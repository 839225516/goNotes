package main

import "fmt"

func main() {
	var breakAgain bool

	// 外循环
	for x := 0; x < 10; x++ {

		// 内循环
		for y := 0; y < 10; y++ {

			// 满足某条件时退出循环
			if y == 2 {

				// 设置退出标记
				breakAgain = true

				// 退出内循环
				break
			}
		}

		// 根据标记，还需要退出一次循环
		if breakAgain {
			break
		}
	}

	fmt.Println("done")

}
