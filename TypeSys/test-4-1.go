package main

import "fmt"

type X struct {
	a int
}

type Y struct {
	X
	b int
}

type Z struct {
	Y
	c int
}

func (x X) Print() {
	fmt.Printf("In X, a = %d\n", x.a)
}

func (y Y) Print() {
	fmt.Printf("In Y, b = %d\n", y.b)
}

func (z Z) Print() {
	fmt.Printf("In Z, c = %d\n", z.c)
}

func main() {
	x := X{a: 1}

	y := Y{
		X: x,
		b: 2,
	}

	z := Z{
		Y: y,
		c: 3,
	}

	println(z.a, z.Y.a, z.Y.X.a) // 指的是同一个

	z.Print() // 输出 In Z, c = 3
	x.Print() // 输出 In X, a = 1
}
