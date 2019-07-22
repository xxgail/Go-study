package main

import (
	"fmt"
)

/**
Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。
*/

// 定义结构体
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

// 该 method 属于 Circle 类型对象中的方法
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
func (c Circle) getArea() float64 {
	// c.radius 即为Circle类型对象中的属性
	return 3.14 * c.radius * c.radius
}

type FangXing struct {
	a float64
}

func (f FangXing) getArea() float64 {
	return f.a * f.a
}
