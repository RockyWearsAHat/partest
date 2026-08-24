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
		{"cjk_longer", "日本語", "語本日"},
		{"emoji", "😀😁", "😁😀"},
		{"mixed_script", "hello世界", "界世olleh"},
		{"empty", "", ""},
		{"single_char", "a", "a"},
		{"single_rune_cjk", "本", "本"},
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
		{"accents/cafe", "café", "CAFÉ"},
		{"accents/naive", "naïve", "NAÏVE"},
		{"already_upper", "HELLO", "HELLO"},
		{"mixed_case", "HeLLo", "HELLO"},
		{"empty", "", ""},
		{"single_char", "a", "A"},
		{"single_accent", "é", "É"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Upper(tt.s); got != tt.want {
				t.Errorf("Upper(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
