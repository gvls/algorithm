package sort

// package sort

// QuickSort 挖坑 + 分治 + 递归
func QuickSort(src []int, lef int, rig int) {
	if lef >= rig {
		return
	}

	i := lef
	j := rig
	tmp := src[i] // hole is left
	for i < j {
		for i < j && tmp <= src[j] { // left  <=  right
			j--
		}
		if i < j {
			src[i] = src[j] // hole is right
		}

		for i < j && src[i] <= tmp { // left  =>  right
			i++
		}
		if i < j {
			src[j] = src[i] // hole is left
		}
	}
	src[i] = tmp

	QuickSort(src, lef, i-1) // divide left part
	QuickSort(src, i+1, rig) // divide right part
}
