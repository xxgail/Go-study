package main

import (
	"fmt"
)

func main() {
	var a = "Gail"
	//var b = "Hello" + "," + a + "!"
	b := "Hello"
	fmt.Println(a, b)
	aaa := roundArea(1.0)
	fmt.Print(aaa)

	//sum := 0
	//for i:= 0; i<10; i++  {
	//	sum += 1
	//}
	//fmt.Print(sum)

	arr1 := map[string]int{"a": 1, "b": 2}

	for k, v := range arr1 {
		fmt.Printf("the key is %v and the value is %v \n", k, v)
	}
}

func roundArea(a float32) float32 {
	const Pi = 3.14

	xx := a * a * Pi

	return xx
}
