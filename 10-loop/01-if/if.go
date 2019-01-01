package main

import (
	"fmt"
	"os"
)

func main() {
	/* if 的格式如下：
	 if 表达式1 {
		分支1
	} else if 表达式2 {
		分支2
	} else {
		分支3
	}
	*/

	var num int = 11
	if num > 10 {
		fmt.Println("num > 10")
	} else {
		fmt.Println("num <= 10")
	}

	// if 的特殊写法，可以在 if 的表达式之前添加一个执行语句
	// os.Open()是一个有返回值的函数
	// err != nil 才是 if 的判断表达式
	if _, err := os.Open("test"); err != nil {
		fmt.Println(err)
		return
	}

}
