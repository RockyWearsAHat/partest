package mathx

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		// normal cases (both positive)
		{"positives/2+3", 2, 3, 5},
		{"positives/5+10", 5, 10, 15},
		// negative operands
		{"negatives/2+(-3)", 2, -3, -1},
		{"negatives/(-2)+3", -2, 3, 1},
		{"negatives/(-2)+(-3)", -2, -3, -5},
		// zero cases
		{"zero/0+0", 0, 0, 0},
		{"zero/0+5", 0, 5, 5},
		{"zero/5+0", 5, 0, 5},
		// same operands
		{"same/5+5", 5, 5, 10},
		{"same/3+3", 3, 3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		// normal cases (both positive)
		{"positives/2-3", 2, 3, -1},
		{"positives/5-3", 5, 3, 2},
		// negative operands
		{"negatives/2-(-3)", 2, -3, 5},
		{"negatives/(-2)-3", -2, 3, -5},
		{"negatives/(-2)-(-3)", -2, -3, 1},
		// zero cases
		{"zero/0-0", 0, 0, 0},
		{"zero/0-5", 0, 5, -5},
		{"zero/5-0", 5, 0, 5},
		// same operands
		{"same/5-5", 5, 5, 0},
		{"same/3-3", 3, 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.a, tt.b); got != tt.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
