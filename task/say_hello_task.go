package task

import (
	"fmt"
	"runtime/debug"
	"time"
)

func Init() {
	Timer(3*time.Second, 3*time.Second, SayHello, "", nil, nil)
}

func SayHello(param interface{}) (result bool) {
	result = true

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SayHello stop", r, string(debug.Stack()))
		}
	}()

	fmt.Println("hello,小盖")

	return
}
