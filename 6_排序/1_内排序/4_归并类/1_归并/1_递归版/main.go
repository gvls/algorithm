package main

import "fmt"

// package main
// O(nlogn) O(n+logn)

// main
func main() {
	a := []int{4, 5, 2, 1, 6, 3, 9, 0, 7, 8}
	dst := []int{4, 5, 2, 1, 6, 3, 9, 0, 7, 8}
	//a := []int{4, 5, 2, 1}
	//tr1 := []int{4, 5, 2, 1}
	fmt.Println(a)
	fmt.Println("")
	MergeSort(a, dst, 0, len(a)-1)
	fmt.Println("")
	fmt.Println(dst)
}

// MergeSort
func MergeSort(src []int, dst []int, b int, e int) {
	var tmp []int = make([]int, 10)
	if b == e {
		dst[b] = src[b]
		return
	}

	m := (b + e) / 2
	MergeSort(src, tmp, b, m)   // src[b]-src[m] -有序-> tmp[b]-tmp[m]
	MergeSort(src, tmp, m+1, e) // src[m+1]-src[e] -有序-> tmp[m+1]-tmp[e]
	merge(tmp, dst, b, m, e)    // tmp[b]-tmp[m] + tmp[m+1]-tmp[e] -> dst
}

// merge
func merge(tmp []int, dst []int, i int, m int, n int) {
	iii := i
	test(tmp, i, m, n)
	poi := i
	j := m + 1
	for ; i <= m && j <= n; poi++ { // 从这两个片段中不断取最小值插入 dst
		if tmp[i] < tmp[j] {
			dst[poi] = tmp[i]
			i++
		} else {
			dst[poi] = tmp[j]
			j++
		}
	}

	if i <= m { // 直接插入剩下的
		for l := 0; l <= m-i; l++ {
			dst[poi+l] = tmp[i+l]
		}
	}

	if j <= n { // 直接插入剩下的
		for l := 0; l <= n-j; l++ {
			dst[poi+l] = tmp[j+l]
		}
	}

	test2(dst, iii, n)
}

// test
func test(tmp []int, i int, m int, n int) {
	for k := i; k <= m; k++ {
		fmt.Print(tmp[k], " ")
	}
	fmt.Print(" +  ")
	for k := m + 1; k <= n; k++ {
		fmt.Print(tmp[k], " ")
	}
	fmt.Print(" => ")
}

// test2
func test2(dst []int, b int, e int) {
	for i := b; i <= e; i++ {
		fmt.Print(dst[i], " ")
	}
	fmt.Println("")
	fmt.Println("")
}
