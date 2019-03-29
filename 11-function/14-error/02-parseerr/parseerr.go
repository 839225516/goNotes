package main

import "fmt"

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {
	var e error

	//创建一个错误实例，包含文件名和行号
	e = newParseError("main.go", 10)

	// 通过 error 接口查看错误描述
	fmt.Println(e.Error())

	// 根据错误的具体类型，获取详细的错误信息
	switch detail := e.(type) {
	case *ParseError: //这是一个解析错误
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
}
