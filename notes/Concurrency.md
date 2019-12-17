# 并发

~~本章内容全部来自于《Go语言核心编程》第五章~~

## 1.并发基础
#### 1-1. 并发和并行
- 并发意味着程序在任意时刻都是同时运行的。
- 并行意味着程序在单位时间内是同时运行的。
#### 1-2.并发执行体
- goroutine: go + 函数（匿名函数或者有名函数）
- 特征
    - go的执行是非阻塞的
    - go后面的函数的返回值会被忽略
    - 所有的goroutine是平等的被调度和执行的
#### 1-3.chan 通道 用于通信
- 参考示例 [test-1-3.go](../Concurrency/test-1-3.go)
- chan 通道是goroutine之间通信的基础，主要指程序的数据通道。
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
- 没有缓冲的通道用于**同步**功能
#### 1-4.WaitGroup
- 参考示例 [test-1-4.go](../Concurrency/test-1-4.go)
- 可以用于实现多个goroutine同步的机制
- WaitGroup 用来等待多个goroutine完成，main goroutine 调用Add设置需要等待goroutine的数目，每一个goroutine结束时调用Done(),Wait()被main用来等待所有的goroutine完成。
#### 1-5. select
- 参考示例 [test-1-5.go](../Concurrency/test-1-5.go)
- **多路监听多个通道**。如果通道没有可读或可写的，就堵塞。如果有一个，就进入处理。如果有多个，就随机选一个处理。
#### 1-6. 扇入和扇出
#### 1-7. 通知退出机制
- 参考示例 [test-1-7.go](../Concurrency/test-1-7.go)
- 关闭select监听的某个通道能使select立即感知这种通知

## 2.并发范式
#### 2-1. 生成器 
- 参考示例 [test-2-1.go](../Concurrency/test-2-1.go)
- 常见的就是调用一个统一全局的生成器服务，用于生成全局事务号、订单号、序列号和随机数等。
#### 2-2. 管道
- 参考示例 [test-2-2.go](../Concurrency/test-2-2.go)
- 通道可以分为两个方向，一个是读一个是写。输入类型和输出类型相同的话，函数就可以调用自己，形成一个调用链。
#### 2-3. 每个请求一个goroutine
- 参考示例 [test-2-3.go](../Concurrency/test-2-3.go)

## 3.Context标准库
#### 3-1. context的设计目的
- 如果遇到程序崩溃或者需要重新启动，goroutine不会自动停止。除非goroutine在函数体内return或者主goroutine终止运行，否则我们是不能通过外部手段干扰goroutine使其终止的。会造成goroutine leak
- 跟踪goroutine调用树，并在这些goroutine调用树中传递通知和元数据。
    - 退出通知机制 ——— 通知可以传递给整个goroutine调用树上的每一个goroutine。
    - 传递数据 ——— 数据可以传递给整个goroutine调用树上的每一个goroutine。
#### 3-2. 基本数据结构
- 参考[context的使用](https://www.cnblogs.com/apocelipes/p/10344011.html)
- context接口
```go
package Context

type Context interface {
    // 返回超时时间（duration加上创建context对象时的时间），如果已经超时ok为true
    // 返回的时间也可以是自己设置的time.Time
    Deadline() (deadline time.Time, ok bool)

    // done信号，和上一节的做法一样，这里进行了一些包装
    Done() <-chan struct{}

    // 如果Done未被关闭就返回nil。
    // 否则返回相应的错误，比如调用了cancel()会返回Canceled；超时会返回DeadlineExceeded
    Err() error

    // 可以给context设置一些值，使用方法和map类似，key需要支持==比较操作，value需要是并发安全的
    Value(key interface{}) interface{}
}
```    
#### 3-3. API函数
- 参考示例 [test-3-3.go](../Concurrency/test-3-3.go)
- context包提供了三种context，分别是是普通context，超时context以及带值的context：
```go
  // 普通context，通常这样调用ctx, cancel := context.WithCancel(context.Background())
  // 带有退出通知的Context具体对象
  func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
  
  // 带超时通知的context，超时之后会自动close对象的Done，与调用CancelFunc的效果一样
  // WithDeadline 明确地设置一个d指定的系统时钟时间，如果超过就触发超时
  // WithTimeout 设置一个相对的超时时间，也就是deadline设为timeout加上当前的系统时间
  // 因为两者事实上都依赖于系统时钟，所以可能存在微小的误差，所以官方不推荐把超时间隔设置得太小
  // 通常这样调用ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
  func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
  
  // 带有值的context，可以传递数据，没有CancelFunc，所以它只用于值的多goroutine传递和共享
  // 通常这样调用ctx := context.WithValue(context.Background(), "key", myValue)
  func WithValue(parent Context, key, val interface{}) Context
```
- 这些函数的共同特点——parent参数。可以逐层包装并传递。
#### 3-4. 辅助函数
