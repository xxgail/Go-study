# 并发

~~本章内容全部来自于《Go语言核心编程》第五章~~

## 1.并发基础
### 1-1. 并发和并行
- 并发意味着程序在任意时刻都是同时运行的。
- 并行意味着程序在单位时间内是同时运行的。
### 1-2.并发执行体
- goroutine: go + 函数（匿名函数或者有名函数）
- 特征
    - go的执行是非阻塞的
    - go后面的函数的返回值会被忽略
    - 所有的goroutine是平等的被调度和执行的
### 1-3.chan 通道
- 参考示例 test-1-3.go
- 通过make来创建，分为有缓冲的通道和无缓冲的
```
// 无缓冲的通道，通道内存放元素的类型为datatype
make(chan datatype )

// 有10个缓冲的通道，通道内存放元素的类型为datatype
make(chan datatype, 10)
```
- 无缓冲的通道len和cap都是0
- 有缓存的通道的len代表没有被读取的元素数、cap代表整个通道的容量。
- 读取已经关闭的通道不会引发阻塞
### 1-4.WaitGroup
### 1-5. select
- 参考示例 test-1-5.go
- **多路监听多个通道**。如果通道没有可读或可写的，就堵塞。如果有一个，就进入处理。如果有多个，就随机选一个处理。
### 1-6. 扇入和扇出
### 1-7. 通知退出机制
- 参考示例 test-1-7.go
- 关闭select监听的某个通道能使select立即感知这种通知

## 2.并发范式
### 2-1. 生成器 
- 参考示例 test-2-1.go
- 常见的就是调用一个统一全局的生成器服务，用于生成全局事务号、订单号、序列号和随机数等。
### 2-2. 管道
- 参考示例 test-2-2.go
- 通道可以分为两个方向，一个是读一个是写。输入类型和输出类型相同的话，函数就可以调用自己，形成一个调用链。
### 2-3. 每个请求一个goroutine
- 参考示例 test-2-3.go