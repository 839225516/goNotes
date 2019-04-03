package main

import "fmt"

/*
	接口被实现的条件二： 接口中所有方法均被实现
*/

type DataWriter interface {
	WriteData(data interface{}) error

	// 能否写入
	CanWrite() bool
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("write data: ", data)
	return nil
}

func (d file) CanWrite() bool {
	fmt.Println("can Write")
	return true
}

func main() {
	// 实例化 file
	f := new(file)

	var writer DataWriter

	writer = f

	writer.CanWrite()
}
