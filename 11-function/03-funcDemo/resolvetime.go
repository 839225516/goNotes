package main

import "fmt"

const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60

	// 定义每小时的秒数
	SecondsPreHour = SecondsPerMinute * 60

	// 定义每天的秒数
	SecondsPerDay = SecondsPerMinute * 12
)

// 将传入的秒数 解析为 分，时，天
func resolveTime(seconds int) (day int, hour int, minute int) {

	day = seconds / SecondsPerDay
	hour = seconds / SecondsPreHour
	minute = seconds / SecondsPerMinute

	return
}

func main() {
	// fmt.Println() 使用了可变参数，可以接收不定量的参数
	fmt.Println(resolveTime(1000))

	// 只获取小时和分钟
	_, hour, minute := resolveTime(18000)
	fmt.Println(hour, minute)

	// 只获取天数
	day, _, _ := resolveTime(19000)
	fmt.Println(day)
}
