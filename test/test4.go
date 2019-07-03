package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 定义类型。定义funcType为一个 “没有参数，返回值为int类型的函数”
type funcType func() int

func fib() funcType {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (f funcType) Read(p []byte) (n int, err error) {
	next := f()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d", next)
	return strings.NewReader(s).Read(p)
}

func scan(read io.Reader) {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf(text)
	}
}

func main() {
	f := fib()
	scan(f)
}
