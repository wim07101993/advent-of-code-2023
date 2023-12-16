package main

import (
	"advent-of-code-2023/shared"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	input := shared.ReadAllRunesByLine(strings.NewReader(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`))
	expected := shared.ReadAllRunesByLine(strings.NewReader(`....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`))

	output := Expand(input)

	if len(output) != len(expected) {
		t.Errorf("expected %v lines but got %v", len(expected), len(output))
		return
	}
	for y := range expected {
		if len(output[y]) != len(expected[y]) {
			t.Errorf("expected %v characters but got %v", len(expected[y]), len(output[y]))
			continue
		}
		for x := range expected[y] {
			if output[y][x] != expected[y][x] {
				t.Errorf("expected character at %v,%v to be %v but got %v", x, y, expected[y][x], output[y][x])
			}
		}
	}
}

func TestSolve(t *testing.T) {
	cases := []struct {
		input    string
		expected int
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
			expected: 374,
		},
	}

	for _, c := range cases {
		output := Solve(strings.NewReader(c.input))
		if output != c.expected {
			t.Errorf("expected output to be %v but got %v", c.expected, output)
		}
	}
}
