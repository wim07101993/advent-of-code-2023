package main

import (
	shared10 "advent-of-code-2023/day10/shared"
	"advent-of-code-2023/shared"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	rs := shared.ReadAllRunesByLine(r)

	sx, sy := shared10.FindStartingPoint(rs)
	l := CalcLoopLength(sx, sy, rs)

	return l / 2
}

func CalcLoopLength(sx, sy int, rs [][]rune) int {
	const north = "north"
	const south = "south"
	const west = "west"
	const east = "east"

	x, y := sx, sy
	comeFrom := ""

	length := 0
	for {
		length++

		if comeFrom != north && shared10.CanGoNorth(x, y, rs) {
			comeFrom = south
			y--
		} else if comeFrom != south && shared10.CanGoSouth(x, y, rs) {
			comeFrom = north
			y++
		} else if comeFrom != east && shared10.CanGoEast(x, y, rs) {
			comeFrom = west
			x++
		} else if comeFrom != west && shared10.CanGoWest(x, y, rs) {
			comeFrom = east
			x--
		} else {
			panic("Dead end")
		}

		if rs[y][x] == 'S' {
			return length
		}

	}
}
