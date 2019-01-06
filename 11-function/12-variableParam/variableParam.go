package main

import (
	"bytes"
	"fmt"
)

// 可变参数
/*
	func 函数名(固定参数列表, v ... T)(返回参数列表){
		函数体
	}
*/

func joinStrings(slist ...string) string {
	//定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer

	// 遍历可变参数列表 slist ,类型为 []string
	for _, s := range slist {
		// 将遍历的字符串连续写入字节数组
		b.WriteString(s)
	}

	// 将连接好的字节数组转换为字符串
	return b.String()
}

func main() {
	fmt.Println(joinStrings("pig ", "and", "rat"))
	fmt.Println(joinStrings("hammer", " mom"))
}
