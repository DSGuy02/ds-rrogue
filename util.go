// Code written by Oladeji Sanyaolu (util.go) 08/11/2023

package main

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
