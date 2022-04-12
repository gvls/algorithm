package main

import "fmt"

// package main 不断找出最小的 放在排序好的序列后面 (小 -> 大)
// O(n^2) 性能大于冒泡

// Select
func Select(a []int) {
	for i := 0; i < len(a)-1; i++ { // count of select max
		fmt.Println(a)
		min := i
		for j := i + 1; j < len(a); j++ { // select current max in remaining
			if a[min] > a[j] { // 因为可能会找到多次，所以使用min，让多次交换变成一次交换
				min = j
			}
		}
		if i != min { // when search already, swap
			a[i], a[min] = a[min], a[i]
		}
	}
	fmt.Println(a)
}

// main
func main() {
	a := []int{4, 5, 2, 1, 6, 3, 9, 0, 7, 8}
	Select(a)
}
