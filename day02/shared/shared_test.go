package shared

import (
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
				Id: 1,
				Rounds: []Round{
					{Blue: 3, Red: 4},
					{Red: 1, Green: 2, Blue: 6},
					{Green: 2},
				},
			},
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Game{
				2,
				[]Round{
					{Blue: 1, Green: 2},
					{Green: 3, Blue: 4, Red: 1},
					{Green: 1, Blue: 1},
				},
			},
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Game{
				3,
				[]Round{
					{Green: 8, Blue: 6, Red: 20},
					{Blue: 5, Red: 4, Green: 13},
					{Green: 5, Red: 1},
				},
			},
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Game{
				4,
				[]Round{
					{Green: 1, Red: 3, Blue: 6},
					{Green: 3, Red: 6},
					{Green: 3, Blue: 15, Red: 14},
				},
			},
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Game{
				5,
				[]Round{
					{Red: 6, Blue: 1, Green: 3},
					{Blue: 2, Red: 1, Green: 2},
				},
			},
		},
	}

	success := true
	for _, c := range cases {
		output := ParseGame(c.input)
		if !expectGameEquals(t, output, c.expected) {
			success = false
		}
	}
	if !success {
		t.Fail()
	}
}

func expectGameEquals(t *testing.T, received Game, expected Game) bool {
	success := true
	if received.Id != expected.Id {
		t.Logf("expected game to have Id %v but got %v", expected.Id, received.Id)
		success = false
	}
	if len(received.Rounds) != len(expected.Rounds) {
		t.Logf("expected game to have %v Rounds but got %v", len(expected.Rounds), len(received.Rounds))
		success = false
	}

	for i := range received.Rounds {
		if !expectRoundEquals(t, received.Rounds[i], expected.Rounds[i]) {
			success = false
		}
	}
	return success
}

func expectRoundEquals(t *testing.T, received Round, expected Round) bool {
	success := true
	if received.Red != expected.Red {
		t.Logf("expected round to have %v Red dice but got %v", expected.Red, received.Red)
		success = false
	}
	if received.Green != expected.Green {
		t.Logf("expected round to have %v Red dice but got %v", expected.Green, received.Green)
		success = false
	}
	if received.Blue != expected.Blue {
		t.Logf("expected round to have %v Red dice but got %v", expected.Blue, received.Blue)
		success = false
	}
	return success
}
