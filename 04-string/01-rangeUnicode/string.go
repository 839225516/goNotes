package main

import "fmt"

func main() {
	theme := "你好，hello"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii :  %c  %d\n", theme[i], theme[i])
	}

	// for range 使用unicode 遍历字符串
	for _, s := range theme {
		fmt.Printf("unicode: %c  %d\n", s, s)
	}
}
