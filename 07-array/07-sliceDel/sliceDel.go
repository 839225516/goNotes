package main

import "fmt"

func main() {

	// go 本身没有直接提供删除切片元素的方法，需要append()函数来连接切片

	seq := []string{"a", "b", "c", "d", "e"}

	// 指定删除的下标索引
	index := 2

	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])

	// 将删除点的前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)

	fmt.Println(seq)

}
