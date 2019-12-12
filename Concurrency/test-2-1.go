package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA2(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Label:
		for {
			select {
			case ch <- <-GenerateIntA2(send):
			case ch <- <-GenerateIntB(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	// 创建一个作为接收退出信号的chan
	done := make(chan struct{})

	// 启动生成器
	ch := GenerateInt(done)

	// 获取生成器资源
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// 通知生产者停止生产
	done <- struct{}{}
	fmt.Println("Stop generate")
}
