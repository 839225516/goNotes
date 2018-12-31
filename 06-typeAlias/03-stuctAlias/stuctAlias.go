package main

import (
	"fmt"
	"reflect"
)

// 定义商标结构
type Brand struct {
}

// 为商标结构添加 Show()方法
func (t Brand) Show() {

}

// 为 Brand 定义一个别名 FakeBrand
type FackBrand = Brand

//定义车辆结构
type Vehicle struct {
	// 嵌入两个结构体
	// 并不意味着嵌入了两个Brand, FackBrand的类型会以名字的方式保留在Vechicle成员中
	Brand
	FackBrand
}

func main() {
	// 声明变量 a 为车辆类型
	var a Vehicle

	// 指定调用FackBrand的show方法
	a.FackBrand.Show()

	// 取a的类型反射对象
	ta := reflect.TypeOf(a)

	//遍历 a 的所有成员
	for i := 0; i < ta.NumField(); i++ {
		//a的成员的信息
		f := ta.Field(i)

		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}

}
