package main

import "fmt"

//将 NewInt 定义为 int 类型
// NewInt 会形成新的类型，本身依然具备int的特性
type NewInt int

// 将int 取一个别名叫 IntAlias
// IntAlias 和 int 等效 ，且只会在代码中存在，编译完成后，不会有IntAlias类型
type IntAlias = int

func main() {

	var a NewInt

	fmt.Printf("a type: %T\n", a)

	var b IntAlias
	fmt.Printf("b type: %T\n", b)

}
