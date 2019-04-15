package main

import (
	"fmt"
	"sync/atomic"
)

var (
	//序列号
	seq int64
)

func GenID() int64 {
	// 尝试原子的增加序列号
	atomic.AddInt64(&seq, 1)
	return seq

	// 修改解决竞态问题
	// return atomic.AddInt64(&seq, 1)
}

func main() {

	// 循环10次GenID,同时忽略GenID的返回值
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID)
}

// 运行时参数加入 "-race" 参数，开启运行时对竞态问题的分析
// go run -race racedetect.go
