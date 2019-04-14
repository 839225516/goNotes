package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func server(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	fmt.Println("listen: " + address)

	defer l.Close()

	// 侦听循环
	for {
		// 新连接没有到来时，accept是阻塞的
		conn, err := l.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// 根据连接开启会话，这个过程需要并行执行
		go handleSession(conn, exitChan)
	}
}

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("session started:")

	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')

		// 数据读取正确
		if err == nil {
			// 去掉字符串尾部的回车
			str = strings.TrimSpace(str)

			// 处理Telnet指令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}

			// Echo逻辑，发什么数据，原样返回
			conn.Write([]byte(str + "\r\n"))
		} else {
			//发生错误
			fmt.Println("session closed")
			conn.Close()
			break
		}
	}
}

// telnet命令处理
func processTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("server shutdown")

		exitChan <- 0

		return false
	}

	fmt.Println(str)

	return true
}

func main() {
	exitChan := make(chan int)

	go server("127.0.0.1:7001", exitChan)

	code := <-exitChan

	os.Exit(code)
}
