package main

import (
	day11 "advent-of-code-2023/day11/shared"
	"advent-of-code-2023/shared"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(Solve(os.Stdin, 1_000_000))
}

func Solve(r io.Reader, scale int64) int64 {
	rs := shared.ReadAllRunesByLine(r)
	gs := day11.ListGalaxies(rs)
	shrunkCols := GetShrunkCols(rs)
	shrunkRows := GetShrunkRows(rs)

	var sum int64
	for i := range gs {
		for j := i + 1; j < len(gs); j++ {
			x := ShortestDistanceBetween(gs[i], gs[j], shrunkCols, shrunkRows, scale)
			sum += x
		}
	}
	return sum
}

func ShortestDistanceBetween(c1, c2 day11.Coordinate, shrunkCols, shrunkRows []int, shrinkScale int64) int64 {
	dist := int64(day11.AbsDiff(c1.X, c2.X) + day11.AbsDiff(c1.Y, c2.Y))
	for _, x := range shrunkCols {
		if c1.X > c2.X {
			if x > c2.X && x < c1.X {
				dist += shrinkScale - 1
			}
		} else {
			if x < c2.X && x > c1.X {
				dist += shrinkScale - 1
			}
		}
	}
	for _, y := range shrunkRows {
		if c1.Y > c2.Y {
			if y > c2.Y && y < c1.Y {
				dist += shrinkScale - 1
			}
		} else {
			if y < c2.Y && y > c1.Y {
				dist += shrinkScale - 1
			}
		}
	}
	return dist
}

func GetShrunkRows(rs [][]rune) []int {
	var ret []int
	for y := range rs {
		if day11.IsEmptyRow(rs, y) {
			ret = append(ret, y)
		}
	}
	return ret
}

func GetShrunkCols(rs [][]rune) []int {
	var ret []int
	for x := range rs[0] {
		if day11.IsEmptyColumn(rs, x) {
			ret = append(ret, x)
		}
	}
	return ret
}
