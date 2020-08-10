package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有goroutine都要增加其值的变量
	counter int

	// wg 原来等待程序结束
	wg sync.WaitGroup

	// mutex 用来定义一段代码临界区
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter

			// 当前goroutine从线程退出，并放回到队列
			runtime.Gosched()

			value++

			counter = value
		}
		// 释放锁，允许其他正在等待的goroutine进入临界区
		mutex.Unlock()
	}
}
