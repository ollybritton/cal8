package cal8

import "testing"

// TestCenter tests that the Center utility works.
func TestCenter(t *testing.T) {
	tests := []struct {
		in, expected string
		size         int
	}{
		{"January", "       January      ", 20},
	}

	for _, tc := range tests {
		got := Center(tc.in, " ", tc.size)
		if got != tc.expected {
			t.Errorf("fail, got '%s', expecting '%s'", got, tc.expected)
		}
	}
}

// TestHorizontalAppend tests that two strings can be horizontally appended.
func TestHorizontalAppend(t *testing.T) {
	tests := []struct {
		s1, s2, expected string
	}{
		{`hello
how
are
you`, `i
am
good
thanks`, `hello i
how   am
are   good
you   thanks`},
		{`hello
how
are
you`, `i
missing
one`, `hello i
how   missing
are   one
you   `},
		{`i
missing
one`, `hello
how
are
you`, `i       hello
missing how
one     are
        you`},
	}

	for _, tc := range tests {
		got := HorizontalAppend(tc.s1, tc.s2, " ")
		if got != tc.expected {
			t.Errorf("fail, got '%s', expecting '%s'", got, tc.expected)
		}
	}
}
