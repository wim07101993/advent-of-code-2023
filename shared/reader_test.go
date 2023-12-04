package shared

import (
	"strings"
	"testing"
)

func TestReadAllRunesByLine(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	expected := [][]rune{
		[]rune("467..114.."),
		[]rune("...*......"),
		[]rune("..35..633."),
		[]rune("......#..."),
		[]rune("617*......"),
		[]rune(".....+.58."),
		[]rune("..592....."),
		[]rune("......755."),
		[]rune("...$.*...."),
		[]rune(".664.598.."),
	}

	output := ReadAllRunesByLine(strings.NewReader(input))

	if len(output) != len(expected) {
		t.Fatalf("expected %v lines, got %v", len(expected), len(output))
	}
	for i := range output {
		if string(output[i]) != string(expected[i]) {
			t.Fatalf("expected line %v to be %v, got %v", i, string(expected[i]), string(output[i]))
		}
	}
}
