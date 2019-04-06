package main

import "io"

/*
	go 中 接口与接口也可以嵌套也新的接口
	eg: io包中的 Writer , Closer, WriteCloser

	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	type Closer interface {
		Closer() error
	}

	type WriteCloser interface {
		Writer
		Closer
	}
*/

// 声明一个设备结构
type device struct {
}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

// 实现io.Closer的Close()方法
func (d *device) Close() error {
	return nil
}

func main() {
	//声明写入关闭器，并赋予device的实例
	var wc io.WriteCloser = new(device)

	// 关闭设备
	wc.Close()

	// 声明写入器，并赋予device的新实例
	var writeOnly io.Writer = new(device)

	// 写入数据
	writeOnly.Write(nil)
}
