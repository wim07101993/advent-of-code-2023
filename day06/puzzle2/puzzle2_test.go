package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	output := Solve(strings.NewReader(input))
	const expected = 71503
	if output != expected {
		t.Errorf("expected output to be %v but got %v", expected, output)
	}
}
