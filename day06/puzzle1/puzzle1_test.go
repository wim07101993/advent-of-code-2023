package main

import (
	"strings"
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

func TestSolve(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	output := Solve(strings.NewReader(input))
	const expected = 288
	if output != expected {
		t.Errorf("expected output to be %v but got %v", expected, output)
	}
}
