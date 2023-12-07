package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestParseHandType(t *testing.T) {
	cases := []struct {
		input    string
		expected HandType
	}{
		{"32T4K", HighCard},
		{"32T3K", OnePair},
		{"T55J5", FourOfAKind},
		{"T5535", ThreeOfAKind},
		{"KK677", TwoPair},
		{"KTJJT", FourOfAKind},
		{"QQQJA", FourOfAKind},
		{"QQQJQ", FiveOfAKind},
		{"QQQQQ", FiveOfAKind},
		{"QQ22Q", FullHouse},
	}

	for _, c := range cases {
		output := ParseHandType(c.input)
		if output != c.expected {
			t.Errorf("expected %v to parse to %v but got %v", c.input, c.expected, output)
		}
	}
}

func TestHands_Sort(t *testing.T) {
	input := Hands{
		ParseHand("32T3K 0"),
		ParseHand("T55J5 0"),
		ParseHand("KK677 0"),
		ParseHand("KTJJT 0"),
		ParseHand("QQQJA 0"),
		ParseHand("65K2J 973"),
		ParseHand("JKTTJ 766"),
	}
	rand.Shuffle(len(input), func(i, j int) {
		input[i], input[j] = input[j], input[i]
	})
	expected := Hands{
		ParseHand("32T3K 0"),
		ParseHand("65K2J 973"),
		ParseHand("KK677 0"),
		ParseHand("JKTTJ 766"),
		ParseHand("T55J5 0"),
		ParseHand("QQQJA 0"),
		ParseHand("KTJJT 0"),
	}

	input.Sort()

	for i := range expected {
		if input[i].Strength != expected[i].Strength {
			t.Errorf("Did not expect %v in position %v", input[i].Strength, i)
		}
	}
}

func TestSolve(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	output := Solve(strings.NewReader(input))
	expected := 5905
	if output != expected {
		t.Errorf("expected output to be %v but got %v", expected, output)
	}
}
