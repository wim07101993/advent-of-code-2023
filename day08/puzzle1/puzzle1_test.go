package main

import (
	"strings"
	"testing"
)

func TestNodes_CalcSteps(t *testing.T) {
	cases := []struct {
		nodes    Nodes
		input    string
		expected int
	}{
		{
			Nodes{
				"AAA": {"BBB", "CCC"},
				"BBB": {"DDD", "EEE"},
				"CCC": {"ZZZ", "GGG"},
				"DDD": {"DDD", "DDD"},
				"EEE": {"EEE", "EEE"},
				"GGG": {"GGG", "GGG"},
				"ZZZ": {"ZZZ", "ZZZ"},
			},
			"RL",
			2,
		},
		{
			Nodes{
				"AAA": {"BBB", "BBB"},
				"BBB": {"AAA", "ZZZ"},
				"ZZZ": {"ZZZ", "ZZZ"},
			},
			"LLR",
			6,
		},
	}

	for _, c := range cases {
		output := c.nodes.CalcSteps(c.input)
		if output != c.expected {
			t.Fatalf("expected output to be %v but got %v", c.expected, output)
		}
	}
}

func TestSolve(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			2,
		},
		{
			`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			6,
		},
	}

	for _, c := range cases {
		output := Solve(strings.NewReader(c.input))
		if output != c.expected {
			t.Fatalf("expected output to be %v but got %v", c.expected, output)
		}
	}
}
