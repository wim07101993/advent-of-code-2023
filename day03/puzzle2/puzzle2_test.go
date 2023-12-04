package main

import (
	"advent-of-code-2023/shared"
	"slices"
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
	const expected = 467835
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
		expected int
	}{
		{
			GetNumberInput{[]rune("467..114.."), 0},
			467,
		},
		{
			GetNumberInput{[]rune("467..114.."), 6},
			114,
		},
		{
			GetNumberInput{[]rune(".......+.58"), 10},
			58,
		},
	}

	for _, c := range cases {
		output := getNumber(c.input.rs, c.input.start)
		if output != c.expected {
			t.Fatalf("expected number to be %v but got %v", c.expected, output)
		}
	}
}

type GetAdjacentNumbersCase struct {
	x        int
	y        int
	expected []int
}

func TestGetAdjacentNumbers(t *testing.T) {
	s := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......7555
...$.*....
..664.598.`
	rs := shared.ReadAllRunesByLine(strings.NewReader(s))
	cases := []GetAdjacentNumbersCase{
		{3, 1, []int{467, 35}},
		{3, 4, []int{617}},
		{5, 8, []int{7555, 664, 598}},
	}

	for _, c := range cases {
		output := getAdjacentNumbers(rs, c.x, c.y)
		if len(output) != len(c.expected) {
			t.Fatalf("expected %v but got %v (lenght difference)", c.expected, output)
		}
		for _, i := range c.expected {
			if !slices.Contains(output, i) {
				t.Fatalf("expected %v but got %v (doesnt contain %v)", c.expected, output, i)
			}
		}
	}
}
