package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// 超时控制终止goroutine
//func coroutine(ctx context.Context, duration time.Duration, id int, wg *sync.WaitGroup)  {
//	for  {
//		select {
//		case <- ctx.Done():
//			fmt.Printf("goroutine %d finish\n",id)
//			wg.Done()
//			return
//		case <- time.After(duration):
//			fmt.Printf("message from goroutine %d\n",id)
//		}
//	}
//}
//
//func main()  {
//	wg := &sync.WaitGroup{}
//	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
//	defer cancel()
//
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go coroutine(ctx,1 * time.Second, i, wg)
//	}
//
//	wg.Wait()
//}

/**
使用WaitGroup等待所有的goroutine执行完毕，在收到<-ctx.Done()的终止信号后使wg中需要等待的goroutine数量减一。
因为context只负责取消goroutine，不负责等待goroutine运行，所以需要配合一点辅助手段。
*/

// =======================================================

// 用普通的context取消goroutine

//func main()  {
//	// gen是一个生成器，返回从1开始的递增数字直到自身被取消
//	gen := func(ctx context.Context) <-chan int {
//		dst := make(chan int)
//		n := 1
//		go func() {
//			for  {
//				select {
//				case <- ctx.Done():
//					return
//				case dst <- n:
//					n++
//				}
//			}
//		}()
//		return dst
//	}
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	for n := range gen(ctx) {
//		fmt.Println(n)
//		// 生成到5时循环终止
//		if n == 5 {
//			break
//		}
//	}
//}

// ===================================================

// 在goroutine间共享变量

func main() {
	var v int64
	wg := sync.WaitGroup{}
	ctx := context.WithValue(context.Background(), "myKey", &v)

	for i := 0; i < 10; i++ { // 有10个goroutine分别对v原子地加了一。
		wg.Add(1)
		go func(ctx context.Context, key string) {
			// 取出来是interface{}，需要先断言成我们需要的类型
			value := ctx.Value(key).(*int64)

			atomic.AddInt64(value, 1)
			wg.Done()
		}(ctx, "myKey")
	}

	wg.Wait()

	// 类型断言成*int64然后解引用
	fmt.Println(*(ctx.Value("myKey").(*int64)))
}
