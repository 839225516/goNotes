package main

import (
	"fmt"
	"reflect"
)

// 使用 reflect.TypeOf() 可以获得任意值的类型对象

func main() {
	var a int

	typeOfA := reflect.TypeOf(a)

	// 获取变量的类型名和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
