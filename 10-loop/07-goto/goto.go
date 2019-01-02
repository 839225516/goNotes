package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {

			if y == 2 {

				// 跳转到标签
				goto breakHere
			}
		}
	}

	// 手动返回，避免执行进入标签
	return

	// 标签只能被goto使用，但不影响代码的执行
	// 如果不手动返回，在不满足退出循环条件时，也会执行下面的语句
breakHere:
	fmt.Println("done")

}
