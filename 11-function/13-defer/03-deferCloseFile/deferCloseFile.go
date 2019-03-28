package main

import (
	"fmt"
	"os"
)

// 根据文件名查询文件大小
func fileSzie(filename string) int64 {
	f, err := os.Open(filename)

	// 如果打开文件时出错，返回文件大小为0
	if err != nil {
		return 0
	}

	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时出错，关闭文件并返回文件大小为0
	if err != nil {
		defer f.Close()
		return 0
	}

	// 取文件大小
	size := info.Size()

	// 关闭文件
	f.Close()

	return size
}

func main() {
	file := "deferCloseFile.go"

	fsize := fileSzie(file)

	fmt.Println(fsize)
}
