package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input :=
		`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	const expected = 2

	output := Solve(strings.NewReader(input))
	if output != expected {
		t.Fatalf("expected output to be %v but got %v", expected, output)
	}
}
