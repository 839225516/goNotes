package main

import (
	"fmt"
	"strings"
)

func main() {
	tracer := "死神来了，死神 bye bye"
	comma := strings.Index(tracer, "，")

	// strings.Index 对中文字符串只能匹配到第一个字
	pos := strings.Index(tracer[comma:], "死神")
	posp := strings.Index(tracer[comma:], "死")

	fmt.Printf("%d %d\n", pos, posp)
	fmt.Printf("%c\n", tracer[pos])
	fmt.Printf("%c\n", tracer[posp])
	fmt.Println(tracer[pos:])
	fmt.Println(comma, pos, tracer[comma+pos:])
}
