package pkg1

import (
	"14-package/04-anonymous/pkg2"
	"fmt"
)

/*
	init() 函数的特性
	每个源码文件可以使用1个init()函数
	init()函数会在main()执行前被自动调用
	调用顺序为main中引用的，以深度优先顺序初始化
	  如包的引用关系为 main -> A -> B -> C; 那这些包的init()函数调用顺序为：
		C.init() -> B.init() -> A.init() -> main
	同一个包中的多个init()函数的调用顺序不可预期
	init() 函数不能被其它函数调用
*/

func ExecPkg1() {
	fmt.Println("ExecPkg1")
	pkg2.ExecPkg2()
}

func init() {
	fmt.Println("pkg1 init")
}
