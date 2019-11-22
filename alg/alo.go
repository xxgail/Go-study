package main

import (
	"fmt"
	"github.com/xxgail/alg/AdvantageCount"
	"sort"
)

func main() {
	a := []int{2, 0, 4, 1, 2}
	b := []int{1, 3, 0, 0, 2}
	ad := AdvantageCount.AdvantangeCount(a, b)
	fmt.Print(ad)
	//eq := EightQueue.EightQueue{
	//	Column: make([]int, 8),
	//}
	//eq.CalEightQueues(0)
	//data := heightChecker([]int{1, 1, 4, 2, 1, 3})
	//fmt.Println(data)

	//data := constructArray(7,2)
	//fmt.Println(data)
}

/** 667
*给定两个整数 n 和 k，你需要实现一个数组，这个数组包含从 1 到 n 的 n 个不同整数，同时满足以下条件：
*① 如果这个数组是 [a1, a2, a3, ... , an] ，那么数组 [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数；.
*② 如果存在多种答案，你只需实现并返回其中任意一种.
 */
func constructArray(n int, k int) []int {
	data := []int{1}
	if k == 1 {
		for i := 1; i <= n; i++ {
			data = append(data, i)
		}
	} else {
		for i := 1; i <= k; i++ {
			if i%2 == 0 {
				data = append(data, (i+2)/2)
			} else {
				data = append(data, 1+k-(i-2)/2)
			}
		}

		for i := k + 1; i < n; i++ {
			data = append(data, i+1)
		}
	}
	return data
}

/** 1051
* 学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。
* 请你返回至少有多少个学生没有站在正确位置数量。
* 该人数指的是：能让所有学生以 非递减 高度排列的必要移动人数。
 */
func heightChecker(heights []int) int {
	NewHeight := make([]int, 0, len(heights)/2)

	a := 0
	for i := 0; i < len(heights); i++ {
		NewHeight = append(NewHeight, heights[i])
	}
	sort.Ints(NewHeight)
	for i := 0; i < len(heights); i++ {
		if heights[i]^NewHeight[i] != 0 {
			a = a + 1
		}
	}
	return a
}
