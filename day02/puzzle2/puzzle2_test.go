package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := strings.Builder{}
	input.WriteString("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n")
	input.WriteString("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n")
	input.WriteString("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n")
	input.WriteString("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n")
	input.WriteString("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green\n")

	output := Solve(strings.NewReader(input.String()))
	const expected = 2286
	if output != expected {
		t.Errorf("expected output to be %v but got %v", expected, output)
	}
}
