package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{"IV"}, 4},
		{"9", args{"IX"}, 9},
		{"58", args{"LVIII"}, 58},
		{"1994", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
			if got := romanToIntNew(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_romanToInt(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			romanToInt("MCMXCIV")
		}
	})

	b.Run("byte", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			romanToIntNew("MCMXCIV")
		}
	})
}
