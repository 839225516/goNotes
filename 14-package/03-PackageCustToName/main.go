package main

// 在导入包路径前添加标识符即可形成自定义引用包
// customName "path/to/package"

import (
	// 自定义引用包名
	renameLib "14-package/03-PackageCustToName/mylib"
	"fmt"
)

func main() {
	// fmt.Println(mylib.Add(1,3)
	fmt.Println(renameLib.Add(1, 3))
}
