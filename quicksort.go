package leetcode

import (
	"math/rand"
)

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

// qsort3 is a 3 way quick sort solution
func qsort3(slice []int) {
	n := len(slice)
	if n < 2 {
		return
	}

	swap := func(a, b int) {
		slice[a], slice[b] = slice[b], slice[a]
	}

	// pick and move the pivot to the left
	pivotIdx := rand.Intn(n)
	pivotV := slice[pivotIdx]
	swap(0, pivotIdx)

	l, r := 0, n
	for i := 1; i < r; i++ {
		if slice[i] < pivotV {
			// move elements smaller than pivot to the left
			swap(i, l+1)
			l++
		} else if slice[i] > pivotV {
			// move elements bigger than pivot to the right
			swap(i, r-1)
			r--
			i--
		}
	}

	// move the pivot(s) after the last smaller element
	swap(0, l)

	// Recursion
	qsort3(slice[:l])
	qsort3(slice[r:])
}
