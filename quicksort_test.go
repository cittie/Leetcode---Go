package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{5, 1}, []int{1, 5}},
		{[]int{5, 5, 5}, []int{5, 5, 5}},
		{[]int{2, 1, 14, 25, 7, 2444}, []int{1, 2, 7, 14, 25, 2444}},
		{
			[]int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 5, 22, 211, 211, 2211, 22233, 33321, 2312, 132, 546, 77567, 7675, 767, 88, 88, 88, 88, 88, 88, 45646, 74564, 464564, 4564564, 648867},
			[]int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 5, 22, 88, 88, 88, 88, 88, 88, 132, 211, 211, 546, 767, 2211, 2312, 7675, 22233, 33321, 45646, 74564, 77567, 464564, 648867, 4564564},
		},
	}

	for _, test := range tests {
		qsort(test.in)
		assert.Equal(t, test.out, test.in)
	}

	for _, test := range tests {
		qsort3(test.in)
		assert.Equal(t, test.out, test.in)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	src := []int{
		1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 5, 22, 211, 211, 2211, 22233, 33321, 2312, 132, 546, 77567, 7675, 767, 88, 88, 88, 88, 88, 88, 45646, 74564, 464564, 4564564, 648867,
	}

	b.Run("quick sort 2 way", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			qsort(src)
		}
	})

	b.Run("quick sort 3 way", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			qsort3(src)
		}
	})
}
