package main

import "fmt"

func main() {

	// 宕机前defer语句会优先被执行
	defer fmt.Println("宕机后要做的事情 1")
	defer fmt.Println("宕机后要做的事情 2")

	panic("宕机")
}
