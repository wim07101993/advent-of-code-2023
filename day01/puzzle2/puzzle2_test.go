package main

import (
	"strings"
	"testing"
)

func TestSolveLine(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, c := range cases {
		output := SolveLine(c.input)
		if output != c.expected {
			t.Logf("expected output to be %v but got %v", c.expected, output)
		}
	}
}

func TestSolve(t *testing.T) {
	input := strings.Builder{}
	input.WriteString("two1nine\n")
	input.WriteString("eightwothree\n")
	input.WriteString("abcone2threexyz\n")
	input.WriteString("xtwone3four\n")
	input.WriteString("4nineeightseven2\n")
	input.WriteString("zoneight234\n")
	input.WriteString("7pqrstsixteen\n")

	output := Solve(strings.NewReader(input.String()))
	const expected = 281
	if output != expected {
		t.Logf("expected output to be %v but got %v", expected, output)
	}
}
