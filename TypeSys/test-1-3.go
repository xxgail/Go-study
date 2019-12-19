package main

import "fmt"

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "gaga"

	var ma Map = mp // 因为有相同底层类型，所以可以赋值。而且mp是未命名类型。

	//var im iMap = ma // 没有一个未命名类型，所以不能赋值

	ma.Print()

	// Map 实现了 Print()，所以其可以赋值给接口类型变量
	var i interface {
		Print()
	} = ma

	i.Print()
}
