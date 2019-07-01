package main

import "fmt"

/**
接口
*/
type Phone interface {
	call() // 定义接口里的方法
}

type NokiaPhone struct {
}

func (np NokiaPhone) call() {
	fmt.Println("I am Nokia")
}

type IPhone struct {
}

func (ip IPhone) call() {
	fmt.Println("I am iPhone")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
