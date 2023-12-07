package shared

import (
	"testing"
)

func TestMagic(t *testing.T) {
	cases := []struct {
		t        int
		d        int
		expected int
	}{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
	}

	for _, c := range cases {
		output := Magic(c.t, c.d)
		if output != c.expected {
			t.Errorf("expected t %v and d %v to yield %v but got %v", c.t, c.d, c.expected, output)
		}
	}
}
