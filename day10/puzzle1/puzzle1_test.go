package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			input: `.....
.S-7.
.|.|.
.L-J.
.....`,
			expected: 4,
		},
		{
			input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`,
			expected: 8,
		},
	}

	for _, c := range cases {
		output := Solve(strings.NewReader(c.input))
		if output != c.expected {
			t.Errorf("expected output to be %v but got %v", c.expected, output)
		}
	}
}
