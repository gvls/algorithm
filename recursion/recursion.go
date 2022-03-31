package recursion

// package recursion

// Recursion self call self ...
func Recursion(key int) { // have different variate in each call
	// exit
	if key == 0 {
		return
	}

	// do somethings and tendency to exit
	key--
	Recursion(key) // call own
}
