package main

import "fmt"

// channel 在多个goroutine间通信的管道
// go 提倡使用通道channel 代替共享共存

/*
	channel 特性
	channel 是一种特殊的类型，
	在从任何时候，同时只能有一个goroutine访问通道进行发送和获取数据
	goroutine 间通过channel就可以通信
	channel 像一个队列，遵循先入先出的规则，保证收发数据的顺序

	// 声明通道的类型
	var 通道变量 chan 通道类型
		通道类型： 通道内的数据类型
		通道变量： 保存通道的变量

	chan类型的空值是 nil, 声明后需要配合make后才能使用

	// 创建通道
	通道实例 := make(chan 数据类型)
	ch1 := make(chan int)
	ch2 := make(chan interface{})
	type Equip struct{...}
	ch3 := make(chan *Equip)


	// 使用通道
	// 创建一个空接口通道
	ch := make(chan interface{})

	// 将0放入通道中
	ch <- 0

	// 将hello 放入通道中
	ch <- "hello"

	把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。

	// 使用通道接收数据
	通道接收有如下特性：
		通道的收发操作在不同的两个goroutine间进行
		接收将持续阻塞直到发送方发送数据
		每次接收一个元素

	1) 阻塞接收数据
	data  := <- ch

	2) 非阻塞接收数据,语句不会发生阻塞
	data, ok := ch
	ok 表示是否接收到数据
	非阻塞的通道接收方法 可能造成高的CPU占用，因此使用非常少

	3）接收任意数据，忽略接收的数据
	<- ch
*/

func main() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")

		// 通过通道通知main的goroutine
		ch <- 0

		fmt.Println("exit goroutine")
	}()

	fmt.Print("wait goroutine")

	// 等待匿名 goroutine
	<-ch

	fmt.Println("all done")
}
