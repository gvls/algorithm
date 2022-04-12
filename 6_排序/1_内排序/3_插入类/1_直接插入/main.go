package main

import "fmt"

// package main 不断将记录插入 排好序的序列中
// O(n^2) 性能大于选择排序

// Insert
func Insert(a []int) {
	for i := 1; i < len(a); i++ { // sort remaining
		fmt.Println(a)
		newData := a[i]                // 要插入的 新记录
		j := i - 1                     // 排好序的序列的 末尾记录
		for 0 <= j && newData < a[j] { // 向前遍历，search place to Insert
			a[j+1] = a[j] // 不断 move left to right
			j--
		}
		// Insert
		if j+1 != i { // 插入位置是否是 原位置
			a[j+1] = newData // Insert
		}
	}
	fmt.Println(a)
}

// Insert2
func Insert2(a []int) {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] { // 假如已经排好序了，就不用再插入了
			fmt.Println(a)
			newData := a[i]
			j := i - 1
			for 0 <= j && newData < a[j] {
				a[j+1] = a[j]
				j--
			}
			a[j+1] = newData
		}
	}
	fmt.Println(a)
}

// main
func main() {
	a := []int{4, 5, 2, 1, 6, 3, 9, 0, 7, 8}
	Insert2(a)
}
