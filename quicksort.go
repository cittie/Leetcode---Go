package leetcode

import "math/rand"

// qsort is a 2 way quick sort solution
func qsort(slice []int) {
	n := len(slice)
	if n < 2 {
		return
	}

	swap := func(a, b int) {
		slice[a], slice[b] = slice[b], slice[a]
	}

	// pick and move the pivot to the right
	pivotIdx := rand.Intn(n)
	swap(pivotIdx, n-1)

	// move elements smaller than pivot to the left
	left := 0
	for i := range slice {
		if slice[i] < slice[n-1] {
			swap(i, left)
			left++
		}
	}

	// move the pivot after the last smaller element
	swap(left, n-1)

	// Recursion
	qsort(slice[:left])
	qsort(slice[left+1:])
}
