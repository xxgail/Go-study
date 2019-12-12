package main

import (
	"fmt"
	"sync"
)

// 工作任务
type task struct {
	begin  int
	end    int
	result chan<- int
}

// 任务执行：计算从begin到end的和
// 执行结果写入结果chan result
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func main() {
	// 创建任务通道
	taskchan := make(chan task, 10)

	// 创建结果通道
	resultchan := make(chan int, 10)

	// wait 用于同步等待任务的执行
	wait := &sync.WaitGroup{}

	// 初始化task的goroutine，计算100个自然数之和
	go InitTask(taskchan, resultchan, 100)

	// 每个task启动一个goroutine进行处理
	go DistributeTask(taskchan, wait, resultchan)

	// 通过结果通道获取结果并汇总
	sum := ProcessResult(resultchan)

	fmt.Println("sum=", sum)
}

// 构建task并写入task通道
func InitTask(taskchan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskchan <- tsk
	}

	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskchan <- tsk
	}

	close(taskchan)
}

// 读取task chan ，每个task启动一个worker goroutine 进行处理
// 并等待每个task运行完，关闭结果通道
func DistributeTask(taskchan <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskchan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(result)
}

// goroutine 处理具体工作，并将处理结果发送到结果通道
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

// 读取结果通道，汇总结果
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}

/********************************************************
程序的逻辑分析：
1. InitTask 函数构建task并发送到task通道中
2. 分发任务函数DistributeTask为每个task启动一个goroutine处理任务，等待其处理完成，然后关闭结果通道。
3. ProcessResult函数读取并统计所有结果
*/
