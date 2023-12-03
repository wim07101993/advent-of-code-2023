package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	output := Solve(strings.NewReader(input))
	const expected = 4361
	if output != expected {
		t.Fatalf("expected output to be %v but got %v", expected, output)
	}
}

type GetNumberInput struct {
	rs    []rune
	start int
}

func TestGetNumber(t *testing.T) {
	cases := []struct {
		input    GetNumberInput
		expected string
	}{
		{
			GetNumberInput{[]rune("467..114.."), 0},
			"467",
		},
		{
			GetNumberInput{[]rune("467..114.."), 5},
			"114",
		},
		{
			GetNumberInput{[]rune(".......+.58"), 9},
			"58",
		},
	}

	for _, c := range cases {
		output := getNumber(c.input.rs, c.input.start)
		if output != c.expected {
			t.Fatalf("expected number to be %v but got %v", c.expected, output)
		}
	}
}

func TestReadAllRunesByLine(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	expected := [][]rune{
		[]rune("467..114.."),
		[]rune("...*......"),
		[]rune("..35..633."),
		[]rune("......#..."),
		[]rune("617*......"),
		[]rune(".....+.58."),
		[]rune("..592....."),
		[]rune("......755."),
		[]rune("...$.*...."),
		[]rune(".664.598.."),
	}

	output := readAllRunesByLine(strings.NewReader(input))

	if len(output) != len(expected) {
		t.Fatalf("expected %v lines, got %v", len(expected), len(output))
	}
	for i := range output {
		if string(output[i]) != string(expected[i]) {
			t.Fatalf("expected line %v to be %v, got %v", i, string(expected[i]), string(output[i]))
		}
	}
}

type CheckNumberHasAdjacentSymbolsCase struct {
	s        string
	x        int
	y        int
	expected bool
}

func TestCheckNumberHasAdjacentSymbols(t *testing.T) {
	s := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......7555
...$.*....
.664.598..`
	rs := readAllRunesByLine(strings.NewReader(s))
	cases := []CheckNumberHasAdjacentSymbolsCase{
		{"467", 0, 0, true},
		{"114", 5, 0, false},
		{"35", 2, 2, true},
		{"633", 6, 2, true},
		{"617", 0, 4, true},
		{"58", 8, 5, false},
		{"592", 2, 6, true},
		{"7555", 6, 7, true},
		{"664", 1, 9, true},
		{"598", 5, 9, true},
	}

	for _, c := range cases {
		output := checkNumberHasAdjacentSymbols(rs, c.s, c.x, c.y)
		if output && !c.expected {
			t.Fatalf("expected %v to not have adjacent symbols", c.s)
		} else if !output && c.expected {
			t.Fatalf("expected %v to have adjacent symbols", c.s)
		}
	}
}
