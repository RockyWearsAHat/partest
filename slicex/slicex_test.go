package slicex

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want int
	}{
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"positives", []int{1, 2, 3}, 6},
		{"negatives", []int{-1, -2, -3}, -6},
		{"mixed_sign", []int{5, -3, 2, -1}, 3},
		{"all_zeros", []int{0, 0, 0}, 0},
		{"single_zero", []int{0}, 0},
		{"large_values", []int{1000000, 2000000}, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.xs); got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.xs, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		max  int
		ok   bool
	}{
		{"empty", []int{}, 0, false},
		{"single_positive", []int{5}, 5, true},
		{"single_negative", []int{-3}, -3, true},
		{"positives", []int{1, 5, 3}, 5, true},
		{"negatives", []int{-5, -1, -3}, -1, true},
		{"mixed", []int{-5, 2, -3, 8, 1}, 8, true},
		{"duplicates", []int{5, 3, 5}, 5, true},
		{"all_zeros", []int{0, 0, 0}, 0, true},
		{"single_zero", []int{0}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			max, ok := Max(tt.xs)
			if max != tt.max || ok != tt.ok {
				t.Errorf("Max(%v) = (%d, %v), want (%d, %v)", tt.xs, max, ok, tt.max, tt.ok)
			}
		})
	}
}
