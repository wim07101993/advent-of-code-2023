package main

import (
	day11 "advent-of-code-2023/day11/shared"
	"advent-of-code-2023/shared"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	rs := shared.ReadAllRunesByLine(r)
	rs = Expand(rs)
	gs := day11.ListGalaxies(rs)

	sum := 0
	for i := range gs {
		for j := i + 1; j < len(gs); j++ {
			x := ShortestDistanceBetween(gs[i], gs[j])
			sum += x
		}
	}
	return sum
}

func ShortestDistanceBetween(c1, c2 day11.Coordinate) int {
	return day11.AbsDiff(c1.X, c2.X) + day11.AbsDiff(c1.Y, c2.Y)
}

func Expand(rs [][]rune) [][]rune {
	for y := 0; y < len(rs); y++ {
		if day11.IsEmptyRow(rs, y) {
			rs = slices.Insert(rs, y, rs[y])
			y++
		}
	}

	for x := 0; x < len(rs[0]); x++ {
		if day11.IsEmptyColumn(rs, x) {
			DuplicateColumn(rs, x)
			x++
		}
	}

	return rs
}

func DuplicateColumn(rs [][]rune, x int) {
	for y := 0; y < len(rs); y++ {
		rs[y] = slices.Insert(rs[y], x, rs[y][x])
	}
}
