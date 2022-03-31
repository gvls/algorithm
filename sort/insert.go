package sort

import "fmt"

// package sort

// Insert
func Insert() {
	for i := 1; i < len(a); i++ { // sort remaining
		fmt.Println(a)
		current := a[i]
		poi := i - 1
		for 0 <= poi && a[poi] > current { // search place to Insert
			a[poi+1] = a[poi] // move left to right
			poi--
		}
		// if poi + 1 != i {
		a[poi+1] = current // Insert
		// }
	}
	fmt.Println(a)
}
