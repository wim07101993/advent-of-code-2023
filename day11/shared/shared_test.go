package shared

import (
	"advent-of-code-2023/shared"
	"strings"
	"testing"
)

func TestListGalaxies(t *testing.T) {
	input := shared.ReadAllRunesByLine(strings.NewReader(`....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`))
	expected := []Coordinate{
		{4, 0},
		{9, 1},
		{0, 2},
		{8, 5},
		{1, 6},
		{12, 7},
		{9, 10},
		{0, 11},
		{5, 11},
	}
	output := ListGalaxies(input)

	if len(output) != len(expected) {
		t.Errorf("expected %v galaxies but found %v \n%v\n%v", len(expected), len(output), expected, output)
		return
	}
	for i := range output {
		if output[i] != expected[i] {
			t.Errorf("expected %v at %v but found %v", expected[i], i, output[i])
		}
	}
}
