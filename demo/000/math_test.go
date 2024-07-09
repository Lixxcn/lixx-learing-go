package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{1, 2, 3},
		{-1, 5, 4},
		{0, 0, 0},
	}

	for _, test := range tests {
		actual := Add(test.a, test.b)
		if actual != test.expect {
			t.Errorf("Add(%d, %d) = %d; expect %d", test.a, test.b, actual, test.expect)
		}
	}
}
