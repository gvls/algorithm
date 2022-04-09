package main

import (
	"fmt"
	"time"
)

// package main 对层序遍历堆后得到的数组 取出第一个放到末尾， 调整剩下的结点得到 n-1堆，重复上面步骤
// 选择排序的 优化版 : 在选择最小记录的同时，根据比较结果 对其他记录做出调整

// main
func main() {
	a := []int{4, 5, 2, 1, 6, 3, 9, 0, 7, 8}
	fmt.Println(a)
	Heap(a)
	fmt.Println(a)
}

// Heap
func Heap(a []int) {
	// init
	for i := len(a)/2 - 1; 0 <= i; i-- { // i := len(a)/2-1 它们都有孩子结点
		adjust(a, i, len(a)-1) // 调整 n 个结点为 堆
	}
	// order
	for i := len(a) - 1; 0 < i; i-- {
		aaa++
		a[0], a[i] = a[i], a[0] // 取出最大 放到树末尾，也就是排序好的序列的头
		adjust(a, 0, i-1)       // 调成剩下的n-1结点为 新堆
	}
}

// adjust from bottom to top and from right to left 非叶子结点当作根节点构建大顶堆
func adjust(a []int, node int, end int) {
	fmt.Println(a[node], "结点找最大..")
	test(a, node, end)
	tmp := a[node] // cache value of root node
	// 找出根和子树中的最大值，用最大值赋给根，若找到，继续找其子树中比根大的...
	for j := 2*node + 1; j <= end; j = j*2 + 1 { // 遍历 s 结点的子树及所有子孙，找到最大值
		if j < end && a[j] < a[j+1] { // j < m : 有右孩子，a[j+1]不越界
			j++
		}
		if tmp >= a[j] { // 找到 根结点、左子树、右子树中的最大值
			break
		}
		a[node] = a[j] // update value of root node
		node = j       // s point to 空位
	}
	a[node] = tmp // 填充空位
	fmt.Println("找到最大并交换后 ：")
	test(a, node, end)
}

// aaa test
var aaa int = 0

// test
func test(a []int, s int, m int) {
	//fmt.Println(a[s], " ", a[m])
	//fmt.Println(s, " ", m)
	layer := 1
	i := 0
	for ; layer < len(a); i++ {
		layer = layer * 2
	}

	count := 1
	l := 1
	for j := 0; j < len(a); j++ {
		for k := 0; k < (i+1)*2-l*2; k++ {
			fmt.Print(" ")
		}
		for k := 0; k < count; k++ {
			if j >= len(a) {
				break
			}
			if j >= len(a)-aaa {
				fmt.Print("_")
			} else {
				fmt.Print(a[j])
			}
			for p := 0; p < i-l+1; p++ {
				fmt.Print(" ")
			}
			if k+1 < count {
				j++
			}
		}
		fmt.Println("")

		count = count * 2
		l++
	}

	for i := 0; i < len(a)*2+1-aaa*2; i++ {
		fmt.Print(" ")
	}
	for i := 0; i < aaa; i++ {
		fmt.Print("v ")
	}
	fmt.Println("")
	fmt.Println(a)

	fmt.Println("========================")
	time.Sleep(time.Second * 5)
}
