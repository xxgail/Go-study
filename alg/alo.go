package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := []int{2, 0, 4, 1, 2}
	//b := []int{1, 3, 0, 0, 2}
	//ad := AdvantageCount.AdvantangeCount(a, b)
	//fmt.Print(ad)
	//eq := EightQueue.EightQueue{
	//	Column: make([]int, 8),
	//}
	//eq.CalEightQueues(0)
	//data := heightChecker([]int{1, 1, 4, 2, 1, 3})
	//fmt.Println(data)

	//data := constructArray(7,2)
	//fmt.Println(data)

	A := []int{1, 2, 1, 2, 3}
	K := 2
	data := subarrayWithKDistinct(A, K)
	fmt.Println(data)
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

/** 992.K个不同整数的子数组
给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组
（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）
返回 A 中好子数组的数目。

示例 1：
输出：A = [1,2,1,2,3], K = 2
输入：7
解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
*/
func subarrayWithKDistinct(A []int, K int) int {
	if A == nil || len(A) < K {
		return 0
	}

	hash := make([]int, len(A)+1)

	l, count, result, results := 0, 0, 1, 0

	for r := 0; r < len(A); r++ {
		hash[A[r]]++

		if hash[A[r]] == 1 {
			count++
		}

		for hash[A[l]] > 1 || count > K {
			if count > K {
				result = 1
				count--
			} else {
				result++
			}

			hash[A[l]]--
			l++
		}

		if count == K {
			results += result
		}
	}

	return results
}
