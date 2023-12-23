package main

import (
	"slices"
	"strings"
	"testing"
)

func TestUnfoldSprings(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			".#",
			".#?.#?.#?.#?.#",
		},
		{
			"???.###",
			"???.###????.###????.###????.###????.###",
		},
	}

	for _, c := range cases {
		output := UnfoldSprings(c.input)
		if output != c.expected {
			t.Errorf("expected %v but got %v", c.expected, output)
		}
	}
}

func TestUnfoldDamaged(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1},
			[]int{1, 1, 1, 1, 1},
		},
		{
			[]int{1, 1, 3},
			[]int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3},
		},
	}

	for _, c := range cases {
		output := UnfoldDamaged(c.input)
		if len(output) != len(c.expected) {
			t.Errorf("expected %v items but got %v\n%v\n%v", len(c.expected), len(output), c.expected, output)
		}
		for i := range c.expected {
			if output[i] != c.expected[i] {
				t.Errorf("expected %v at %v but got %v\n%v\n%v", c.expected[i], i, output[i], c.expected, output)
			}
		}
	}
}

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

	cs := make([]chan string, len(cases))
	for i, c := range cases {
		cs[i] = make(chan string)
		go func(i int) {
			GetArrangements(c.springs, c.damaged, cs[i])
		}(i)
		var output []string
		for s := range cs[i] {
			output = append(output, s)
		}

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

	const expected = 525152

	output := Solve(strings.NewReader(input))

	if output != expected {
		t.Errorf("expected %v arrangements but got %v", expected, output)
	}
}
