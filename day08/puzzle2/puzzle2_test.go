package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestParseNodes(t *testing.T) {
	input := `11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	n11A := &Node{Name: "11A"}
	n11B := &Node{Name: "11B"}
	n11Z := &Node{Name: "11Z"}
	n22A := &Node{Name: "22A"}
	n22B := &Node{Name: "22B"}
	n22C := &Node{Name: "22C"}
	n22Z := &Node{Name: "22Z"}
	nXXX := &Node{Name: "XXX"}

	n11A.Left, n11A.Right = n11B, nXXX
	n11B.Left, n11B.Right = nXXX, n11Z
	n11Z.Left, n11Z.Right = n11B, nXXX

	n22A.Left, n22A.Right = n22B, nXXX
	n22B.Left, n22B.Right = n22C, n22C
	n22C.Left, n22C.Right = n22Z, n22Z
	n22Z.Left, n22Z.Right = n22B, n22B

	nXXX.Left, nXXX.Right = nXXX, nXXX

	expected := []*Node{n11A, n11B, nXXX, n11Z, n22A, n22B, n22C, n22Z}

	output := ParseNodes(bufio.NewScanner(strings.NewReader(input)))

	if len(output) != len(expected) {
		t.Errorf("expected %v nodes but got %v", len(expected), len(output))
	}

	for i := range output {
		ExpectNodeEquals(t, output[i], expected[i])
	}
}

func TestGetStarters(t *testing.T) {
	n11A := &Node{Name: "11A"}
	n11B := &Node{Name: "11B"}
	n11Z := &Node{Name: "11Z"}
	n22A := &Node{Name: "22A"}
	n22B := &Node{Name: "22B"}
	n22C := &Node{Name: "22C"}
	n22Z := &Node{Name: "22Z"}
	nXXX := &Node{Name: "XXX"}
	input := []*Node{n11A, n11B, nXXX, n11Z, n22A, n22B, n22C, n22Z}
	expected := []*Node{n11A, n22A}

	output := GetStarters(input)

	if len(output) != len(expected) {
		t.Errorf("expected %v starters but got %v, (%v, %v)", len(expected), len(output), expected, output)
	}
	for i := range output {
		if output[i].Name != expected[i].Name {
			t.Errorf("expected node at %v to have name %v but got %v", i, expected[i].Name, output[i].Name)
		}
	}
}

func TestNode_CalcStepsUntilNextZ(t *testing.T) {
	n11A := &Node{Name: "11A"}
	n11B := &Node{Name: "11B"}
	n11Z := &Node{Name: "11Z"}
	n22A := &Node{Name: "22A"}
	n22B := &Node{Name: "22B"}
	n22C := &Node{Name: "22C"}
	n22Z := &Node{Name: "22Z"}
	nXXX := &Node{Name: "XXX"}

	n11A.Left, n11A.Right = n11B, nXXX
	n11B.Left, n11B.Right = nXXX, n11Z
	n11Z.Left, n11Z.Right = n11B, nXXX

	n22A.Left, n22A.Right = n22B, nXXX
	n22B.Left, n22B.Right = n22C, n22C
	n22C.Left, n22C.Right = n22Z, n22Z
	n22Z.Left, n22Z.Right = n22B, n22B

	nXXX.Left, nXXX.Right = nXXX, nXXX

	instr := "LR"

	cases := []struct {
		node          *Node
		instr         string
		expectedSteps int
		expectedEnd   *Node
	}{
		{n11A, instr, 2, n11Z},
		{n22A, instr, 3, n22Z},
	}

	for _, c := range cases {
		steps, newNode := c.node.CalcStepsUntilNextZ(c.instr)
		if steps != c.expectedSteps {
			t.Errorf("expected %v to end in %v steps but got %v", c.node.Name, c.expectedSteps, steps)
		}
		if newNode != c.expectedEnd {
			t.Errorf("expected %v as end but got %v", *c.expectedEnd, *newNode)
		}
	}
}

func TestSolve(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	expected := 6

	output := Solve(strings.NewReader(input))
	if output != expected {
		t.Fatalf("expected output to be %v but got %v", expected, output)
	}
}

func ExpectNodeEquals(t *testing.T, received *Node, expected *Node) {
	if received.Name != expected.Name {
		t.Errorf("expected node with name %v but got %v (%v, %v)", expected.Name, received.Name, *expected, *received)
	}
	if received.Left.Name != expected.Left.Name {
		t.Errorf("expected node with left %v but got %v (%v, %v)", expected.Left.Name, received.Left.Name, *expected, *received)
	}
	if received.Right.Name != expected.Right.Name {
		t.Errorf("expected node with right %v but got %v (%v, %v)", received.Right.Name, expected.Right.Name, *received, *expected)
	}
}
