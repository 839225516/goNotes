package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// len输出字符串的字节长度
	tip := "genji is a ninja"
	fmt.Println(len(tip))

	// golang 字符串都以utf-8格式保存，每个中文件占用3个字节
	next := "忍者"
	fmt.Println(len(next))

	//go中utf-8包中提供了统计uncode字符数量的函数 RuneCountInString()
	fmt.Println(utf8.RuneCountInString(next))
}
