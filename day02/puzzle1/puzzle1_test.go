package main

import (
	"strings"
	"testing"
)

func TestParseGame(t *testing.T) {
	cases := []struct {
		input    string
		expected Game
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Game{
				id: 1,
				rounds: []Round{
					{blue: 3, red: 4},
					{red: 1, green: 2, blue: 6},
					{green: 2},
				},
			},
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Game{
				2,
				[]Round{
					{blue: 1, green: 2},
					{green: 3, blue: 4, red: 1},
					{green: 1, blue: 1},
				},
			},
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Game{
				3,
				[]Round{
					{green: 8, blue: 6, red: 20},
					{blue: 5, red: 4, green: 13},
					{green: 5, red: 1},
				},
			},
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Game{
				4,
				[]Round{
					{green: 1, red: 3, blue: 6},
					{green: 3, red: 6},
					{green: 3, blue: 15, red: 14},
				},
			},
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Game{
				5,
				[]Round{
					{red: 6, blue: 1, green: 3},
					{blue: 2, red: 1, green: 2},
				},
			},
		},
	}

	for _, c := range cases {
		output := ParseGame(c.input)
		expectGameEquals(t, output, c.expected)
	}
}

func TestSolve(t *testing.T) {
	input := strings.Builder{}
	input.WriteString("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n")
	input.WriteString("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n")
	input.WriteString("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n")
	input.WriteString("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n")
	input.WriteString("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green\n")

	output := Solve(strings.NewReader(input.String()))
	const expected = 8
	if output != expected {
		t.Logf("expected output to be %v but got %v", expected, output)
	}
}

func expectGameEquals(t *testing.T, received Game, expected Game) {
	if received.id != expected.id {
		t.Fatalf("expected game to have id %v but got %v", expected.id, received.id)
	}
	if len(received.rounds) != len(expected.rounds) {
		t.Fatalf("expected game to have %v rounds but got %v", len(expected.rounds), len(received.rounds))
	}
	for i := range received.rounds {
		expectRoundEquals(t, received.rounds[i], expected.rounds[i])
	}
}

func expectRoundEquals(t *testing.T, received Round, expected Round) {
	if received.red != expected.red {
		t.Fatalf("expected round to have %v red dice but got %v", expected.red, received.red)
	}
	if received.green != expected.green {
		t.Fatalf("expected round to have %v red dice but got %v", expected.green, received.green)
	}
	if received.blue != expected.blue {
		t.Fatalf("expected round to have %v red dice but got %v", expected.blue, received.blue)
	}
}
