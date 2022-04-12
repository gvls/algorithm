package main

import (
	"fmt"
	"time"
)

// package main compare all 相邻的2个
// O(n^2)

// Bubbling
func Bubbling(a []int) {
	for i := 0; i < len(a)-1; i++ { // count which do Bubbling
		fmt.Printf("第 %d 轮冒泡\n", len(a)-1-i)
		for j := 0; j < len(a)-i-1; j++ { // move current max to right
			test(a, j, i, false)
			if a[j] > a[j+1] { // compare all 相邻的2个
				a[j], a[j+1] = a[j+1], a[j]
				test(a, j, i, true)
			}
		}
		fmt.Println("")
	}
	fmt.Println(a)
}

// Bubbling
func Bubbling2(a []int) {
	ifFinish := true                            // 标志算法是否已经完成
	for i := 0; i < len(a)-1 && ifFinish; i++ { // true 未完成，继续排序
		ifFinish = false // 初始化为 完成
		fmt.Printf("第 %d 轮冒泡\n", len(a)-1-i)
		for j := 0; j < len(a)-i-1; j++ {
			test(a, j, i, false)
			if a[j] > a[j+1] {
				ifFinish = true // 有数据交换，未完成
				a[j], a[j+1] = a[j+1], a[j]
				test(a, j, i, true)
			}
		}
		fmt.Println("")
	}
	fmt.Println(a)
}

// main
func main() {
	a := []int{4, 5, 2, 1, 6, 3, 9, 0, 7, 8}
	Bubbling2(a)
}

var ttt bool = true

// test
func test(a []int, j int, i int, f bool) {
	fmt.Println(a)
	for k := 0; k < 2*j+1; k++ {
		fmt.Print(" ")
	}
	if f {
		if ttt {
			ttt = false
		} else {
			ttt = true
		}
	}
	if ttt {
		fmt.Print("x y")
	} else {
		fmt.Print("y x")
	}
	for k := 0; k < len(a)*2-2*j-4-2*i; k++ {
		fmt.Print(" ")
	}
	for k := 0; k < i; k++ {
		fmt.Print(" O")
	}
	fmt.Println("")
	fmt.Println("=======================")
	time.Sleep(time.Second * 1)
}
