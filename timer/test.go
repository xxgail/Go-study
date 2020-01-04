package main

import (
	"fmt"
	"log"
	"time"
)

type OnceCron struct {
	tasks  []*Task       // 任务的队列
	add    chan *Task    // 当遭遇新任务的时候
	remove chan string   // 当遭遇到删除任务的时候
	stop   chan struct{} // 当遇到停止信号的时候
	Logger *log.Logger   // 日志
}

type Job interface {
	Run() // 执行接口
}
type Task struct {
	Job     Job    // 要执行的任务
	Uuid    string // 任务标识，删除时用
	RunTime int64  // 执行时间
	Spacing int64  // 间隔时间
	EndTime int64  // 结束时间
	Number  int64  // 总共要次数
}

func (one *OnceCron) Start() {
	// 初始化的时候加一个一年的长定时器，间隔1小时执行一次
	task := getTaskWithFuncSpacing(3600, time.Now().Add(time.Hour*24*365).Unix(), func() {
		log.Println("It's a Hour timer!")
	})
	one.tasks = append(one.tasks, task)
	go one.run() //
}

func (one *OnceCron) run() {

}

func main() {
	input := make(chan interface{})

	go func() {
		for i := 0; i < 5; i++ {
			input <- i
		}
		input <- "Hello, world"
	}()

	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {
		case msg := <-input:
			fmt.Println(msg)

		case <-t1.C:
			println("5s timer")
			t1.Reset(time.Second * 5)

		case <-t2.C:
			println("10s timer")
			t2.Reset(time.Second * 10)
		}

	}
}
