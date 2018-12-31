package main

import "fmt"

func main() {

	//将int定义为Weapon类型
	type Weapon int

	const (
		// 使用 iota 生成常量，iota 起始值为0
		// 一个const声明内的每一行常量声明，将会自动套用前面的 iota 格式，并自动增加
		Arrow       Weapon = iota //开始生成枚举值，默认为 0
		Shuriken                  //1
		SniperRifle               //2
		Rifle
		Blower
	)

	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)

	//使用枚举类型，并赋值
	var weapon Weapon = Blower //4
	fmt.Println(weapon)

	//使用iota生成幂常量
	const (
		// iota 使用位移操作，每次将上一次的值左移一位
		FlagNone = 1 << iota // 1 << 0  = 1
		FlagRed
		FlagGreen
		FlagBlue
	)

	fmt.Printf("%d %d %d %d\n", FlagNone, FlagRed, FlagGreen, FlagBlue)

	//将枚举值按二进制格式输出
	fmt.Printf("%b %b %b %b\n", FlagNone, FlagRed, FlagGreen, FlagBlue)
}
