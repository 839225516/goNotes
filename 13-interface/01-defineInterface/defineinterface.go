package main

import "fmt"

/*
type 接口类型名 interface{
	方法名1(参数列表1)  返回值列表1
	方法名2(参数列表2)  返回值列表2
	...
}
*/

/*
接口被实现的条件：

条件一： 接口的方法与实现接口的类型方法格式一致

常见几种接口无法实现的报错
1）函数名不一致导致的报错
2）实现接口的方法签名不一致导致的报错


*/

// 定义接口
type DataWriter interface {
	WriterData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriteData() 方法
func (d *file) WriterData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriterData: ", data)
	return nil
}

func main() {
	// 实例化file
	f := new(file)

	// 声明一个DataWriter的接口
	var writer DataWriter

	// 将接口赋值f, 也就是*file类型
	writer = f

	// 使用DataWriter接口进行数据写入
	writer.WriterData("data")
}
