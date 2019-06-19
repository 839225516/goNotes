package main

import (
	"context"
	"fmt"
	"time"
)

// context包允许传递一个"context"到程序，Context如超时或截止日期(deadline)或通道，来指示停止或返回。

func Println(ctx context.Context, a, b int) {
	for {
		fmt.Println(a + b)
		a, b = a+1, b+1
		select {
		case <-ctx.Done():
			fmt.Println("程序结束")
			return
		default:
		}
	}
}

func main() {
	{
		//超时取消
		a := 1
		b := 2
		timeout := 2 * time.Second
		ctxBg := context.Background()
		ctx, _ := context.WithTimeout(ctxBg, timeout)
		fmt.Println(ctx, a, b)

		time.Sleep(2 * time.Second) // 等待时间还会继续输出
	}

	{
		// 手动取消
		a := 1
		b := 2
		ctx, cancelCtx := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			cancelCtx() // 在调用主动取消
		}()

		fmt.Println(ctx, a, b)

		time.Sleep(2 * time.Second)
	}

}
