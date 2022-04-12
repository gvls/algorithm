package main

import (
	"fmt"
	"time"
)

// package main 分割成若干小部分，小部分分别进行插入排序，到整个序列基本有序后（不同部分的相同下标：小的在前，大的在后），对整个数据再分割若干更小部分，重复...。全部进行插入排序
// 分割 ： 把距离某个增量的记录组成一个子序列
// 插入排序的 优化版 : 让待排序的记录较少
// O(nlogn) ~ O(n^2)

// main
func main() {
	//a := []int{4, 5, 2, 1, 6, 3, 9, 7, 8}
	a := []int{7, 5, 2, 1, 8, 3, 6, 4, 9}
	Shell(a)
	fmt.Println(a)
}

// Shell
func Shell(a []int) {
	increment := len(a)
	for {
		// 减少增量
		increment = increment/3 + 1 // 2^(t-k+1)-1 0 <= k <= t <= |log2,(n+1)|
		fmt.Println("增量为 ： ", increment)

		// 对目前增量下，对数据分割若干部分，对各个部分的第 k 个位置使用插入排序
		for j := increment; j < len(a); j++ {
			i := j - increment
			test(i, j, a, false, 0)
			if a[i] > a[j] { // 找到 左 > 右
				tmp := a[j]                                     // 缓存 右 边的值
				for ; i >= 0 && a[i] > tmp; i = i - increment { // 插入排序
					a[i+increment] = a[i] // 右边更新为左边的值
					test(i, i+increment, a, true, tmp)
				}
				a[i+increment] = tmp // 左边更新为右边的值
			}
		}
		if increment == 1 {
			break
		}
	}
}

// test
func test(i int, j int, a []int, f bool, tmp int) {
	if f {
		a[i] = 0
	}

	fmt.Print(" ")
	for k := 1; k < 2*len(a); k++ {
		if k%(2*(j-i)) == 0 && k/(2*(j-i)) != 0 {
			fmt.Print(" ")
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println("")

	fmt.Print(a)
	if f {
		fmt.Println(" 待插入: ", tmp)
	} else {
		fmt.Println("        ")
	}

	for k := 0; k < 2*i+1; k++ {
		fmt.Print(" ")
	}
	fmt.Print("i")
	for k := 0; k < 2*(j-i)-1; k++ {
		if f {
			fmt.Print(">")
		} else {
			fmt.Print(" ")

		}
	}
	fmt.Println("j")

	fmt.Println("==============================")
	time.Sleep(time.Second * 1)
}
