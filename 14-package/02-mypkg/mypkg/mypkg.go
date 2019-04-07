package mypkg

/*
	Go中，要在一个包中引用另一个包里的标识符（如类型、变量、常量等）
	必须首先将被引用的标识符导出，将要导出的标识符的 首字母大写 就可以让引用者访问
*/

// myVar 首字母小写，因此只能在mypkg包内使用
var myVar = 100

// Myconst MyStruct 首字母大小是大写，外部可以访问
const MyConst = 89

type Mystruct struct {

	// 首字母大写的成员，外部可以访问
	ExportedField int

	// 首字母小写的，仅限包内访问
	privateMethod()
}

func Add(a, b int) int {
	return a + b
}
