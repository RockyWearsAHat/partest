package strx

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"ascii", "hello", "olleh"},
		{"cjk", "日本", "本日"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.s); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestUpper(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"ascii", "hello", "HELLO"},
		{"accented", "café", "CAFÉ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Upper(tt.s); got != tt.want {
				t.Errorf("Upper(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
