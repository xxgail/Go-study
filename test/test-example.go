package main

import (
	"fmt"
	"time"
)

func main() {
	// Switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	WhatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		default:
			fmt.Println("I don't know")
		}
	}
	WhatAmI(true)

	fmt.Println("----------------------------------------------------------------")

	// Array
	fmt.Println(bbb)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	fmt.Println("----------------------------------------------------------------")

	// Slice
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c", c)

	fmt.Println("----------------------------------------------------------------")

	// Map
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	fmt.Println("----------------------------------------------------------------")

	// Range
	kvs := map[string]string{"a": "apple", "b": "blue"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
		//0 103
		//1 111
	}
	fmt.Println("----------------------------------------------------------------")

	// Variadic Functions
	sum(1, 2)
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
	fmt.Println("----------------------------------------------------------------")

	// Closures-闭包
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("----------------------------------------------------------------")

	// Pointers
	i := 1
	fmt.Println("init:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	fmt.Println("pointer", &i)

	fmt.Println("----------------------------------------------------------------")

	// Methods
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("peri:", r.peri())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("peri:", rp.peri())
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) peri() int {
	return 2*r.width + 2*r.height
}

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}
