package main

import "runtime"

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()

	// ------------------------------------------ 通道
	//c := make(chan struct{}) // 无缓冲的通道
	//go func(i chan struct{}) {
	//	sum := 0
	//	for i := 0; i < 1000; i++ {
	//		sum += i
	//	}
	//	println(sum)
	//
	//	// 写通道
	//	c <- struct{}{}
	//}(c)
	//
	//println("NumGoroutine =", runtime.NumGoroutine())
	//
	////time.Sleep(5 * time.Second)
	//// 读通道
	//<-c

	c := make(chan struct{})  // 无缓冲的通道，通道存放的元素类型为struct{}
	ci := make(chan int, 100) // 有100个缓冲的通道，通道存放的元素类型为int
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i // 将i写入通道ci
		}
		close(ci)

		// 写通道
		c <- struct{}{}
	}(c, ci)

	println("NumGoroutine =", runtime.NumGoroutine())

	// 读通道c,通过通道进行同步等待
	<-c
	//close(c)

	// runtime.NumGoroutine 返回正在执行和排队的任务总数
	println("NumGoroutine1 =", runtime.NumGoroutine())

	// 通道ci虽然已经关闭但仍能读取
	for v := range ci {
		println(v)
	}
}
