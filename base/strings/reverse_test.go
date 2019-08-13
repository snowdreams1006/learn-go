package strings

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		input,expect string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.input)
		if got != c.expect {
			t.Errorf("Reverse(%q) == %q, want %q", c.input, got, c.expect)
		}
	}
}