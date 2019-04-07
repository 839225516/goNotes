package main

import (
	"14-package/05-factory/base"
	_ "14-package/05-factory/cls1" // 匿名引用cls1包，调用init()自动注册
	_ "14-package/05-factory/cls2"
)

func main() {
	// 根据字符串动态创建一个Class1实例
	c1 := base.Create("Class1")
	c1.Do()

	// 根据字符串动态创建一个Class2实例
	c2 := base.Create("Class2")
	c2.Do()
}
