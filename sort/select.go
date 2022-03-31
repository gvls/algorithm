package sort

import "fmt"

// package sort

// Select
func Select() {
	for i := 0; i < len(a)-1; i++ { // count of select max
		fmt.Println(a)
		for j := i + 1; j < len(a); j++ { // select current max in remaining
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}
