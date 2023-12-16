package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		input    string
		scale    int64
		expected int64
	}{
		{
			input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
			scale:    10,
			expected: 1030,
		},
		{
			input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
			scale:    100,
			expected: 8410,
		},
	}

	for _, c := range cases {
		output := Solve(strings.NewReader(c.input), c.scale)
		if output != c.expected {
			t.Errorf("expected output to be %v but got %v", c.expected, output)
		}
	}
}
