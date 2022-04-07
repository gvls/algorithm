package main

import "fmt"

// package main O(logn)
// 二分查找的分支，

var F [12]int = [12]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

// main
func main() {
	a := [12]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, -1, -1}
	k := 21
	index := FibonacciSearch(a, 0, len(a)-1, k)
	fmt.Println(index)
}

// FibonacciSearch
func FibonacciSearch(a [12]int, b int, e int, k int) int {
	p := 0
	for e > F[p]-1 {
		p++
	}

	for i := e; i < F[p]-1; i++ {
		a[i] = a[e]
	}

	for b <= e {
		mid := b + F[p-1] - 1
		if k == a[mid] {
			if mid <= e {
				return mid
			}
			return e
		}
		if k < a[mid] {
			e = mid - 1
			k = k - 1
			continue
		}
		if k > a[mid] {
			b = mid + 1
			k = k - 2
		}

	}

	return -1
}
