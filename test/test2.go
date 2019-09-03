package main

import "fmt"

/**
接口
*/
type Phone interface {
	call() // 定义接口里的方法
}

type NokiaPhone struct {
	name string
}

func (np NokiaPhone) call() {
	np.name = "Nokia"
	fmt.Println("I am " + np.name)
}

type IPhone struct {
}

func (ip IPhone) call() {
	fmt.Println("I am iPhone")
}

type T struct {
	S string
}

func (t *T) call() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var phone Phone

	phone = NokiaPhone{}
	describe(phone)
	phone.call()

	phone = IPhone{} // new(IPhone) <=> &
	describe(phone)
	phone.call()

}

func describe(i Phone) {
	fmt.Printf("(%v, %T)\n", i, i)
}
