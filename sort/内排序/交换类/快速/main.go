package main

import "fmt"

// package main in once sort, divide as part1 < key < part2. repetition do this for part1 and part2 ...
// performance depend on depth of recursion
// 中枢轴 尽量取代排序中的中间数 : reduce recursion
// if source is what have order, performance is low
// not suit small source, can use insertSort
// O(nlogn)~O(n^2) O(logn)~O(n)

// QuickSort 挖坑 + 分治 + 递归
func QuickSort(src []int, lef int, rig int) {
	if lef >= rig {
		return
	}

	i := lef
	j := rig
	// get 中枢轴
	tmp := src[i] // hole is left. key of performance. can is left, right, random, middle of left, right, right,  ...
	for i < j {
		// j from right to left search key which small than tmp
		for i < j && tmp <= src[j] { // left  <=  right
			j--
		}
		if i < j { // replace rather than swap
			src[i] = src[j] // hole is right
		}

		// i from left to right search key which big than tmp
		for i < j && src[i] <= tmp { // left  =>  right
			i++
		}
		if i < j { // repalce rather than swap
			src[j] = src[i] // hole is left
		}
	}
	// set 中枢轴
	src[i] = tmp

	QuickSort(src, lef, i-1) // divide left part
	QuickSort(src, i+1, rig) // divide right part
}

// QuickSort2 reduce count of recursion
func QuickSort2(src []int, lef int, rig int) {
	for lef < rig { // will do left recursion and right recursion

		i := lef
		j := rig
		// get 中枢轴
		tmp := src[i] // hole is left. key of performance. can is left, right, random, middle of left, right, right,  ...
		for i < j {
			// j from right to left search key which small than tmp
			for i < j && tmp <= src[j] { // left  <=  right
				j--
			}
			if i < j { // replace rather than swap
				src[i] = src[j] // hole is right
			}

			// i from left to right search key which big than tmp
			for i < j && src[i] <= tmp { // left  =>  right
				i++
			}
			if i < j { // repalce rather than swap
				src[j] = src[i] // hole is left
			}
		}
		// set 中枢轴
		src[i] = tmp

		QuickSort(src, lef, i-1) // divide left part
		lef = i + 1
	}
}

// main
func main() {
	a := []int{4, 5, 2, 1, 6, 3, 9, 0, 7, 8, 99}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
