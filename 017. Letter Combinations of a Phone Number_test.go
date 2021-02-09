package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	r := letterCombinationsNew("23")

	sort.Strings(expected)
	sort.Strings(r)

	if !reflect.DeepEqual(expected, r) {
		t.Error("Expected: ", expected, "got: ", r)
	}
}

func Benchmark_letterCombinations(b *testing.B) {
	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			letterCombinations("65536")
		}
	})

	b.Run("new", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			letterCombinationsNew("65536")
		}
	})
}
