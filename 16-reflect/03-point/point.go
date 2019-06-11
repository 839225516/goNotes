package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明一个空结构体
	type cat struct {
	}

	// 创建cat实例
	ins := &cat{}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	// 获取类型的元素
	typeOfCat = typeOfCat.Elem()

	fmt.Printf("element name: %v, element kind: %v", typeOfCat.Name(), typeOfCat.Kind())
}
