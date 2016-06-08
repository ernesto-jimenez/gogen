package strconv

import (
	"testing"
)

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"makeSnakeCase", "make_snake_case"},
		{"StartWithCaps", "start_with_caps"},
		{"double_Underscore", "double_underscore"},
	}

	for _, test := range tests {
		res := SnakeCase(test.in)
		if res != test.out {
			t.Errorf("invalid conversion: SnakeCase(%q) should return %q, returned %q", test.in, test.out, res)
		}
	}
}
