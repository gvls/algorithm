package sort

import "fmt"

// package sort

// Bubbling
func Bubbling() {
	for i := 0; i < len(a)-1; i++ { // count which do Bubbling
		fmt.Println(a)
		for j := 0; j < len(a)-i-1; j++ { // move current max to right
			if a[j] > a[j+1] { // compare all 相邻的2个
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
