package main

import "fmt"

// package main
// 减少递归算法中 栈 的消耗，进一步减少时间和空间
// O(nlogn) O(n)

// main
func main() {
	a := []int{4, 5, 2, 1, 6, 3, 9, 0, 7, 8, 99}
	//a := []int{5, 4}
	MergeSort(a)
	fmt.Println(a)
}

// MergeSort
func MergeSort(a []int) {
	tmp := make([]int, len(a))
	k := 1 // 1 -> 2 -> 4 -> 8 -> 16
	for k < len(a) {
		MergePass(a, tmp, k, len(a)-1)
		k = 2 * k
		MergePass(tmp, a, k, len(a)-1) // 对上一次的结果进行下一轮
		k = 2 * k
	}
}

// MergePass
func MergePass(src []int, dst []int, n int, l int) {
	fmt.Println("aa")
	i := 0
	// n 个元素作为一个有序序列
	// 遍历整个数据源，n 与 n 合并成 2n
	for i <= l-2*n+1 {
		fmt.Println(i, "-", i+n-1, "x", i+n, "-", i+2*n-1)
		merge(src, dst, i, i+n-1, i+2*n-1)
		i = i + 2*n
	}

	// 处理末尾
	if i < l-n+1 { // 右边 缺少部分
		fmt.Println("end ", i, "-", i+n-1, "x", i+n, "-", l)
		merge(src, dst, i, i+n-1, l)
	} else { // 缺少右边
		fmt.Println("end add ", i, "-", l)
		for j := i; j <= l; j++ {
			dst[j] = src[j]
		}
	}
}

// merge tmp[i]~tmp[m] + tmp[m+1]~tmp[n] => dst[i]~dst[n]
func merge(tmp []int, dst []int, i int, m int, n int) {
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
}
