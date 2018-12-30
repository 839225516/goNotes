package main

import (
	"bytes"
	"fmt"
)

func main() {
	hammer := "吃我一锤"
	sickle := "死吧"

	// 声明节点缓冲
	var stringBuilder bytes.Buffer

	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)

	fmt.Println(stringBuilder.String())
}
