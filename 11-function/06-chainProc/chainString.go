package main

import (
	"fmt"
	"strings"
)

func StringProcess(list []string, chain []func(string) string) {
	// 遍历每一个字符串
	for index, str := range list {

		// 第一个需要处理的字符串
		result := str

		// 遍历每一个处理链
		for _, proc := range chain {

			// 输入一个字符串进行处理，返回数据作为下一个处理链的输入
			result = proc(result)
		}

		// 将结果放回切片
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {

	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 待处理的函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	// 处理字符串
	StringProcess(list, chain)

	// 输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}
}
