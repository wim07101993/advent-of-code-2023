package main

import (
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

	cs := make([]chan int64, len(cases))
	for i, c := range cases {
		cs[i] = make(chan int64)
		go func(i int) {
			GetArrangementCounts(c.springs, c.damaged, cs[i])
		}(i)
		var total int64
		for s := range cs[i] {
			total += s
		}

		if int(total) != len(c.expected) {
			t.Errorf("expected %v arrangement but got %v\n%v\n%v", len(c.expected), total, c.springs, c.damaged)
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
