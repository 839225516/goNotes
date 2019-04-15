package main

import (
	"fmt"
	"sync"
)

// 在读多写少时，使用读写互斥锁更高效

var (
	count      int
	countGuard sync.RWMutex
)

func GetCount() int {

	// 获取count数据时，将读写互斥锁标记为读状态，另外一个goroutine并发访问了countGuard
	// 同时也调用了countGuard.RLock()时，并不会发生阻塞
	countGuard.RLock()

	defer countGuard.RUnlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {
	SetCount(1)

	fmt.Println(GetCount())

}
