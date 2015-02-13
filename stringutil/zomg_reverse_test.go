package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct { input, expected string } {
		{ "Hello", "olleH" },
		{ "Привет", "тевирП" },
		{"", ""},
	}

	for _, c := range cases {
		got := Reverse(c.input);
		if got != c.expected {
			t.Errorf("Reverse(%q) == %q, expected %q\n", c.input, got, c.expected);
		}
	}
}

