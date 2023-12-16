package main

import (
	"advent-of-code-2023/shared"
	"fmt"
	"io"
	"os"
	"slices"
)

type Coordinate struct {
	X int
	Y int
}

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	rs := shared.ReadAllRunesByLine(r)
	rs = Expand(rs)
	gs := ListGalaxies(rs)

	sum := 0
	for i := range gs {
		for j := i + 1; j < len(gs); j++ {
			x := ShortestDistanceBetween(gs[i], gs[j])
			sum += x
		}
	}
	return sum
}

func ShortestDistanceBetween(c1, c2 Coordinate) int {
	return AbsDiff(c1.X, c2.X) + AbsDiff(c1.Y, c2.Y)
}

func AbsDiff(i1, i2 int) int {
	if i1 > i2 {
		return i1 - i2
	}
	return i2 - i1
}

func ListGalaxies(rs [][]rune) []Coordinate {
	var gs []Coordinate
	for y := range rs {
		for x := range rs[y] {
			if rs[y][x] == '#' {
				gs = append(gs, Coordinate{x, y})
			}
		}
	}
	return gs
}

func Expand(rs [][]rune) [][]rune {
	for y := 0; y < len(rs); y++ {
		if IsEmptyRow(rs, y) {
			rs = slices.Insert(rs, y, rs[y])
			y++
		}
	}

	for x := 0; x < len(rs[0]); x++ {
		if IsEmptyColumn(rs, x) {
			DuplicateColumn(rs, x)
			x++
		}
	}

	return rs
}

func IsEmptyRow(rs [][]rune, y int) bool {
	for x := range rs[y] {
		if rs[y][x] != rs[y][0] {
			return false
		}
	}
	return true
}

func IsEmptyColumn(rs [][]rune, x int) bool {
	for y := range rs {
		if rs[y][x] != rs[0][x] {
			return false
		}
	}
	return true
}

func DuplicateColumn(rs [][]rune, x int) {
	for y := 0; y < len(rs); y++ {
		rs[y] = slices.Insert(rs[y], x, rs[y][x])
	}
}
