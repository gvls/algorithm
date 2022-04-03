package main

import (
	"fmt"
	"time"
)

// package main

// main
func main() {
	S := "ababababx"
	T := "ababx"
	//S = "aaaabcde"
	//T = "aaaaax"
	//    -00123112
	KMP(S, T)
}

// test
func test(S string, T string, i int, j int) {
	if j != -1 { // 测试用
		for k := 0; k < len(S)-i+j; k++ {
			fmt.Print(" ")
		}
		fmt.Println(S)

		for k := 0; k < len(S); k++ {
			fmt.Print(" ")
		}
		for k := 0; k < j; k++ {
			fmt.Print("|")
		}
		if S[i] == T[j] {
			fmt.Println("|")
		} else {
			fmt.Println("x")
		}

		for k := 0; k < len(S); k++ {
			fmt.Print(" ")
		}
		fmt.Println(T)
		fmt.Println("===========================")
	}
}

// KMP
func KMP(S string, T string) {
	i := 0
	j := 0
	next := calNext2(T)
	fmt.Printf("\n========== next >>>>>>>>>>\nnext = %+v\n========== next <<<<<<<<<<\n\n", next)
	count := 0 // 测试用

	for i < len(S) && j < len(T) { // 朴素模式匹配

		count++ // 测试用
		time.Sleep(time.Second * 1)
		test(S, T, i, j) // 测试用

		if j == -1 || S[i] == T[j] { // j==-1 为KMP添加
			i++
			j++
		} else {
			j = next[j] // KMP添加
		}
	}
	fmt.Printf("\n========== count >>>>>>>>>>\ncount = %+v\n========== count <<<<<<<<<<\n\n", count)
	if j >= len(T) {
		fmt.Println(j)
	} else {
		fmt.Println("error")
	}
}

// calNext
func calNext(T string) []int {
	next := make([]int, len(T))
	next[0] = -1 // T[0]匹配失败不使用next，i和j加1后匹配
	i := 1
	j := 0

	for i < len(next)-1 {
		if j == -1 || T[i] == T[j] { // T的第一个匹配失败 或匹配成功
			i++         // T[i] 后缀
			j++         // T[j] 前缀
			next[i] = j // T的 0 ~ i-1 和T前缀 0 ~ j-1 匹配, T的i匹配错误时匹配T的j
		} else {
			j = next[j] // 匹配失败 使用之前生成的next进行next的KMP使用
		}
	}
	//for k := range next { // 调整是否适应下标
	//next[k]++
	//}
	return next
}

// calNext2
func calNext2(T string) []int {
	next := make([]int, len(T))
	next[0] = -1 // T[0]匹配失败不使用next，i和j加1后匹配
	i := 1
	j := 0

	for i < len(next)-1 {
		if j == -1 || T[i] == T[j] { // T的第一个匹配失败 或匹配成功
			i++ // T[i] 后缀
			j++ // T[j] 前缀
			if T[i] != T[j] {
				next[i] = j // T的 0 ~ i-1 和T前缀 0 ~ j-1 匹配, T的i匹配错误时匹配T的j
			} else {
				next[i] = next[j] // T[i]和T[j]相等，当i处出错，j处也出错，因为出错才用next
			}
		} else {
			j = next[j] // 匹配失败 使用之前生成的next进行next的KMP使用
		}
	}
	//for k := range next { // 调整是否适应下标
	//next[k]++
	//}
	return next
}
