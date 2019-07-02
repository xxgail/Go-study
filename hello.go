package main

import (
	"errors"
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

	//// 变量一般用驼峰命名，但是要注意跟函数名区分开
	//roundArea := RoundArea(1.0)
	//fmt.Print(roundArea)

	//// 普通的for循环
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += 1
	//}
	//fmt.Print(sum)

	//// Go 中的普通数组。初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	//arr1 := [3]int{1, 2, 3}
	//arrr1 := [...]int{1, 2, 5}
	//fmt.Print(arr1, arrr1)
	//// 类似于PHP中的foreach循环
	//for _,v := range arr1{
	//	fmt.Println(v)
	//}
	//arr11 := [2][2]int{{0, 0}, {1, 1}}
	//fmt.Print(arr11)

	//// map是Go中的集合，类似于PHP中有key值的数组，类似于js中的对象
	//map1 := map[string]int{"a": 1, "b": 2}
	//// 用range循环集合或者数组，类似于PHP中的foreach
	//for k, v := range map1 {
	//	fmt.Printf("the key is %v and the value is %v \n", k, v)
	//}
	//val,ok := map1["a"]
	//// 查找是否含有"a"这个键，若有则OK返回true，val返回值。没有返回false，和值类型的默认值（值为int型，返回0
	//fmt.Print(ok,val)
	//
	//// 删除map中的键值对，传key就行
	//delete(map1,"a")

	//// 还可以这样写函数，类似于PHP中的匿名函数,赋值给一个变量
	//getSquareRoot := func(x float64) float64 {
	//	return math.Sqrt(x)
	//}
	//fmt.Println(getSquareRoot(9))

	//// 数组的长度不可改变，但是切片可以。
	//slice := []int {1,2,3} // 创建且赋值，可以省略长度
	////var slice1 = make([]int, 2) // 仅创建（
	//// 用make创建 make([]T, length, capacity)：T指类型、length初始长度、capacity容量
	//arr_slice := [5]int{1,2,3,4,5} // 数组
	//s := arr_slice[1:4] // 切片（可省略前后
	//fmt.Println(slice,s)
	//// len()计算切片的长度、cap()计算切片的容量
	//// 此处的cap为4，即从1开始往后有四个位置
	//fmt.Printf("len=%d cap=%d slice=%v\n",len(s),cap(s),s)
	//// 一个切片在未初始化之前默认为 nil，长度为 0

	//var i int
	//for i = 0; i < 10; i++  {
	//	fmt.Printf("%d\t",fibonacci(i))
	//}

	//fmt.Println(factorial(5))

	//val, err := Sqrt(-1)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(val)
	//}

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

/**
用递归来实现斐波那契数列
*/
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

/**
一个斐波那契数列的生成器
*/
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

/**
用递归来实现阶乘
*/
func factorial(n int) int {
	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}

/*
 返回错误error
*/
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), errors.New("")
}
