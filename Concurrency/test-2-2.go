package main

import "fmt"

// chain 就算一个管道（个人理解）。因为输入类型和输出类型相同
func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if v > 5 {
				out <- 10 + v
			} else {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Println(v)
	}
}
