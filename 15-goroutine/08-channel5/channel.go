package main

import (
	"errors"
	"fmt"
	"time"
)

/*
	通道的多路复用
	同一个通道中可以同时接收和发送数据

	go 提供 select 关键字，可以同时响应多个通道的操作。
	select 的每个case都会对应一个通道的收发过程，当收发完成时，就会触发case中的响应语句。
	多个操作在每次select中挑选一个进行响应

	select {
	case 操作1:
		响应操作1
	case 操作2：
		响应操作2
	...
	default:
		没有操作情况
	}

*/

// 模拟rpc客户端的请求和接收消息封装
func RPCClient(ch chan string, req string) (string, error) {
	// 向服务器发送请求
	ch <- req

	// 等待服务器返回
	select {
	case ack := <-ch:
		return ack, nil // 接收到服务器的返回数据
	case <-time.After(time.Second): // 超时
		return "", errors.New("Time out")
	}
}

// 模拟RPC服务器端收到客户端请求和回应
func RPCServer(ch chan string) {
	for {
		// 接收客户端请求
		data := <-ch

		// 打印接收到的数据
		fmt.Println("server received:", data)

		// 模拟超时
		// time.Sleep(time.Second * 2)

		// 向客户端回应
		ch <- "roger"
	}
}

func main() {
	// 创建一个无缓冲字符串通道
	ch := make(chan string)

	go RPCServer(ch)

	// 客户端请求数据和接收数据
	recv, err := RPCClient(ch, "Hi")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received:", recv)
	}
}
