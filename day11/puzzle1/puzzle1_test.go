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

func TestListGalaxies(t *testing.T) {
	input := shared.ReadAllRunesByLine(strings.NewReader(`....#........
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
	expected := []Coordinate{
		{4, 0},
		{9, 1},
		{0, 2},
		{8, 5},
		{1, 6},
		{12, 7},
		{9, 10},
		{0, 11},
		{5, 11},
	}
	output := ListGalaxies(input)

	if len(output) != len(expected) {
		t.Errorf("expected %v galaxies but found %v \n%v\n%v", len(expected), len(output), expected, output)
		return
	}
	for i := range output {
		if output[i] != expected[i] {
			t.Errorf("expected %v at %v but found %v", expected[i], i, output[i])
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
