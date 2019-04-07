package cls1

import (
	"14-package/05-factory/base"
	"fmt"
)

// 定义类1
type Class1 struct {
}

func (c *Class1) Do() {
	fmt.Println("Class1")
}

func init() {
	// 在启动时注册类1工厂
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}
