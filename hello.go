package main

import (
	"fmt"
	"math"
)

func main() {

	// 命名
	var a = "Hello"
	b := "Gail"
	var c string
	c = "!"
	fmt.Println(a, b, c)

	// 变量一般用驼峰命名，但是要注意跟函数名区分开
	roundArea := RoundArea(1.0)
	fmt.Print(roundArea)

	// 普通的for循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Print(sum)

	// Go 中的普通数组。初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	arr1 := [3]int{1, 2, 3}
	arrr1 := [...]int{1, 2, 5}
	fmt.Print(arr1, arrr1)
	arr11 := [2][2]int{{0, 0}, {1, 1}}
	fmt.Print(arr11)

	// map是Go中的集合，类似于PHP中有key值的数组，类似于js中的对象
	arr2 := map[string]int{"a": 1, "b": 2}
	// 用range循环集合或者数组，类似于PHP中的foreach
	for k, v := range arr2 {
		fmt.Printf("the key is %v and the value is %v \n", k, v)
	}

	// 还可以这样写函数，类似于PHP中的匿名函数,赋值给一个变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

}

/**
 求一个圆的面积。
PS: 关于函数名的命名规范
1. 使用驼峰命名
2. 如果包外不需要访问请用小写开头的函数
3. 如果需要暴露出去给包外访问需要使用大写开头的函数名称
*/
func RoundArea(r float32) float32 {
	const Pi = 3.14 // 常量

	area := r * r * Pi

	return area
}
