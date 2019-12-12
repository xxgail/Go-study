package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// 随机数生成器
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label: // 外层的标签，指定循环开始的地方
		for {
			select {
			case ch <- rand.Int():
			case <-done: // 增加一路监听，对退出通知信号done的监听
				break Label // break 跳出指定循环
			}
		}
		// 收到通知后关闭通道ch
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 发送通知，告诉生产者已经停止生产
	close(done)

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	for v := range ch {
		fmt.Println(v)
	}

	// 生产者已经退出
	println("NumGoroutine =", runtime.NumGoroutine())
}
