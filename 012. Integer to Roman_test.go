package leetcode

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"4", args{4}, "IV"},
		{"9", args{9}, "IX"},
		{"58", args{58}, "LVIII"},
		{"1994", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_intToRoman(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			intToRoman(3999)
		}
	})

	b.Run("byte", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			intToRomanNew(3999)
		}
	})
}
