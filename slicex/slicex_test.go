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
		{"multiple", []int{1, 2, 3}, 6},
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
		{"single", []int{5}, 5, true},
		{"multiple", []int{1, 5, 3}, 5, true},
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
