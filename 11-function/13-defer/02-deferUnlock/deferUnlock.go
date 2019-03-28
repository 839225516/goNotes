package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	valueBykey = make(map[string]int)

	//保证使用映射时的并发安全的互斥锁
	valueBykeyGuard sync.Mutex
)

func readValue(key string) int {
	// 对共享资源加锁
	valueBykeyGuard.Lock()

	//取值
	v := valueBykey[key]

	// 对共享资源解锁
	valueBykeyGuard.Unlock()

	fmt.Printf("get valueBykey[%s] is %d\n", key, v)
	return v
}

func setValue(key string, value int) {
	valueBykeyGuard.Lock()

	defer valueBykeyGuard.Unlock()

	fmt.Printf("set valueBykey[%s]=%d\n", key, value)
	valueBykey[key] = value
}

func main() {
	// 开100个协程并发写map
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100000; j++ {
				setValue(fmt.Sprintf("%d", j), j)
			}
		}()
	}

	// 开一个协程读map
	go func() {
		for j := 0; j < 1000; j++ {
			readValue(fmt.Sprintf("%d", j))
		}
	}()

	time.Sleep(time.Second * 30)
}
