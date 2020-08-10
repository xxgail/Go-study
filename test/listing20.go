package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg1.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wg1.Wait()
}

func player(name string, court chan int) {
	defer wg1.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			close(court)
			return
		}

		fmt.Printf("Player %s Hid %d\n", name, ball)
		ball++

		court <- ball
	}
}
