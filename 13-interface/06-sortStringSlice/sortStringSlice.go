package main

import (
	"fmt"
	"sort"
)

/*
	对于string int 等 go提供了固定模式的封装

	字符串切片的便捷排序 sort包中有一个 StringSlice类型
	type StringSlice []string

	func (p StringSlice) len() int { return len(p)}
	func (p StringSlice) Less(i,j int) bool { return p[i] < p[j]}
	func (p StringSlice) Swap(i,j int) { p[i],p[j] = p[j], p[i]}
	func (p StringSlice) Sort() { Sort(p)}
*/

func main() {
	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
