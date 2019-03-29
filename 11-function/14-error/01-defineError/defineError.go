package main

import (
	"errors"
	"fmt"
)

//定义除数为0的错误
// 自定义错误
// var err = errors.New("this is error")
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {

	// 判断除数为0的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}

	return dividend / divisor, nil
}

func main() {
	fmt.Println(div(1, 0))
	fmt.Println(div(2, 3))
}
