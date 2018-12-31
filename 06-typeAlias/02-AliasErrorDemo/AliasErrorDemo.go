package main

import "time"

// 定义time.Duration 的别名为MyDuration
type MyDuration = time.Duration

//为MyDuration添加一个方法
// 编译会报错，不能在非本地的类型 time.Duration 上定义方法： time.Duration在time包中定义
// time.Duration包与main包不在同一个包中，所以不能为time.Duration类型定义方法
// 解决方法： 将 MyDuration 从别名改为类型
func (m MyDuration) EasySet(a string) {

}

func main() {

}
