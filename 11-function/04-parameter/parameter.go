package main

import "fmt"

type Data struct {
	complax  []int      //测试切片在参数中传递的效果
	instance InnerData  //实例分配的 innerData
	ptr      *InnerData // 将ptr声明为InnerData的指针类型
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	// 输出参数的成员情况
	fmt.Printf("inFunc value: %+v\n", inFunc)

	//打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

func main() {

	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},

		ptr: &InnerData{1},
	}

	// 输入结构的成员情况
	fmt.Printf("in Value: %+v\n", in)

	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)

	out := passByValue(in)

	fmt.Printf("out Value: %+v\n", out)

	fmt.Printf("out ptr: %p\n", &out)
}
