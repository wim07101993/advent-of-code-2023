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
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"4sevenfddxgcvdgx", 44},
	}

	for _, c := range cases {
		output := SolveLine(c.input)
		if output != c.expected {
			t.Fatalf("expected output to be %v but got %v", c.expected, output)
		}
	}
}

func TestSolve(t *testing.T) {
	input := strings.Builder{}
	input.WriteString("1abc2\n")
	input.WriteString("pqr3stu8vwx\n")
	input.WriteString("a1b2c3d4e5f\n")
	input.WriteString("treb7uchet\n")

	output := Solve(strings.NewReader(input.String()))
	const expected = 142
	if output != expected {
		t.Fatalf("expected output to be %v but got %v", expected, output)
	}
}
