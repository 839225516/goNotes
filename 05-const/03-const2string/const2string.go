package main

import "fmt"

//定义芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

//定义 ChipType 类型的方法 String(),返回字符串
// 使用String()方法的ChipType在使用上和普通的常量没有区别
// 当这个类型需要显示字符串时，go会自动寻找String()方法并调用
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU aaaa"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	fmt.Printf("%s %d", CPU, CPU)
}
