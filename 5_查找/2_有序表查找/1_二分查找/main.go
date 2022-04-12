package main

import "fmt"

// package main O(logn)
// 待查找的数据是 有序的
// 不断从 中间那个元素 匹配，失败后选择 左边/右边 重复操作

// main
func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 8
	index := BinarySearch(a, 0, len(a)-1, k)
	fmt.Println(index)
}

// BinarySearch
func BinarySearch(a [10]int, b int, e int, k int) int {
	for b <= e {
		mid := (b + e) / 2 // 折半
		if a[mid] == k {   // 找到
			return mid
		}
		if k < a[mid] { // 从前面部分找
			e = mid - 1
			continue
		}
		if k > a[mid] { // 从后面部分找
			b = mid + 1
			continue
		}
	}
	return -1
}
