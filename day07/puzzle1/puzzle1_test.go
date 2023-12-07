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
		{"T55J5", ThreeOfAKind},
		{"KK677", TwoPair},
		{"KTJJT", TwoPair},
		{"QQQJA", ThreeOfAKind},
		{"QQQJQ", FourOfAKind},
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

func TestHand_IsWeakerThan(t *testing.T) {
	cases := []struct {
		hand1    Hand
		hand2    Hand
		expected bool
	}{
		{
			ParseHand("32T3K 0"),
			ParseHand("T55J5 0"),
			true,
		},
		{
			ParseHand("T55J5 0"),
			ParseHand("KK677 0"),
			false,
		},
		{
			ParseHand("KK677 0"),
			ParseHand("KTJJT 0"),
			false,
		},
		{
			ParseHand("KTJJT 0"),
			ParseHand("QQQJA 0"),
			true,
		},
		{
			ParseHand("T55J5 0"),
			ParseHand("QQQJA 0"),
			true,
		},
	}

	for _, c := range cases {
		if c.hand1.IsWeakerThan(c.hand2) && !c.expected {
			if c.expected {
				t.Errorf("expected %v to be stronger than %v", c.hand1, c.hand2)
			} else {
				t.Errorf("expected %v to be stronger than %v", c.hand2, c.hand1)
			}
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
	}
	rand.Shuffle(len(input), func(i, j int) {
		input[i], input[j] = input[j], input[i]
	})
	expected := Hands{
		ParseHand("32T3K 0"),
		ParseHand("KTJJT 0"),
		ParseHand("KK677 0"),
		ParseHand("T55J5 0"),
		ParseHand("QQQJA 0"),
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
	expected := 6440
	if output != expected {
		t.Errorf("expected output to be %v but got %v", expected, output)
	}
}
