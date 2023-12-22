package main

import (
	"slices"
	"strings"
	"testing"
)

func TestGetArrangements(t *testing.T) {
	cases := []struct {
		springs  string
		damaged  []int
		expected []string
	}{
		{
			"???.###",
			[]int{1, 1, 3},
			[]string{"#.#.###"},
		},
		{
			".??..??...?##.",
			[]int{1, 1, 3},
			[]string{
				".#...#....###.",
				".#....#...###.",
				"..#..#....###.",
				"..#...#...###.",
			},
		},
		{
			"?#?#?#?#?#?#?#?",
			[]int{1, 3, 1, 6},
			[]string{".#.###.#.######"},
		},
		{
			"????.#...#...",
			[]int{4, 1, 1},
			[]string{"####.#...#..."},
		},
		{
			"????.######..#####.",
			[]int{1, 6, 5},
			[]string{
				"#....######..#####.",
				".#...######..#####.",
				"..#..######..#####.",
				"...#.######..#####.",
			},
		},
		{
			"?###????????",
			[]int{3, 2, 1},
			[]string{
				".###.##.#...",
				".###.##..#..",
				".###.##...#.",
				".###.##....#",
				".###..##.#..",
				".###..##..#.",
				".###..##...#",
				".###...##.#.",
				".###...##..#",
				".###....##.#",
			},
		},
		{
			"???#???.#??####?",
			[]int{5, 1, 5},
			[]string{
				"#####...#.#####.",
				"#####...#..#####",
				".#####..#.#####.",
				".#####..#..#####",
				"..#####.#.#####.",
				"..#####.#..#####",
			},
		},
		{
			"???#????.#???????",
			[]int{2, 1, 6},
			[]string{
				"##.#.....######..",
				"..##.#...######..",
				"..##..#..######..",
				"..##...#.######..",
				"..##.....#.######",
				"...##.#..######..",
				"...##..#.######..",
				"...##....#.######",
			},
		},
	}

	for _, c := range cases {
		output := GetArrangements(c.springs, c.damaged)
		if len(output) != len(c.expected) {
			t.Errorf("expected %v arrangement but got %v\n%v\n%v\n%v\n%v", len(c.expected), len(output), c.expected, output, c.springs, c.damaged)
		}
		for _, e := range c.expected {
			if !slices.Contains(output, e) {
				t.Errorf("expected to find %v in arrangements\n%v\n%v\n%v", e, output, c.springs, c.damaged)
			}
		}
	}
}

func TestSolve(t *testing.T) {
	input := `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

	const expected = 21

	output := Solve(strings.NewReader(input))

	if output != expected {
		t.Errorf("expected %v arrangements but got %v", expected, output)
	}
}
