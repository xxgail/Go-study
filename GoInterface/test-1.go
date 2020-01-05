package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() { // 隐式声明
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i) // 输出 (<nil>, *main.T)
	i.M()       // 输出 <nil>

	i = &T{"hello"}
	describe(i) // 输出(&{hello}, *main.T)
	i.M()       // 输出hello
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// --------
// 区别于PHP中的implement接口实现方法，go中不需要任何标识符就可以实现该接口内的方法。
// 同样的，如果定义该接口时没有声明任何的方法及变量，就可以认为该接口实现了所有的方法。
