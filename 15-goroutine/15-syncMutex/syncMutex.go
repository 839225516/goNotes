package main

import (
	"fmt"
	"sync"
)

// 互斥锁保证同时只有一个goroutine访问共享资源

var (
	count      int
	countGuard sync.Mutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()

	defer countGuard.Unlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {
	//可以进行并发安全的设置
	SetCount(1)

	// 可以进行并发安全的访问
	fmt.Println(GetCount())
}
