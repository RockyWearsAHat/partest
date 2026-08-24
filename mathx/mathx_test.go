package mathx

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"with negative", -1, 5, 4},
		{"zeros", 0, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Add(c.a, c.b)
			if got != c.want {
				t.Errorf("Add(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 5, 3, 2},
		{"with negative", 5, -1, 6},
		{"zeros", 0, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Sub(c.a, c.b)
			if got != c.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
			}
		})
	}
}
