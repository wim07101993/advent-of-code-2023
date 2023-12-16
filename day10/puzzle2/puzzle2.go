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
	rs[sy][sx] = GetValueForCoordinate(sx, sy, rs)

	rs = GetGridWithoutUnusedPipes(sx, sy, rs)

	return CalculateNrOfEnclosedTiles(rs)
}

func GetValueForCoordinate(x, y int, rs [][]rune) rune {
	cgNorth := shared10.CanGoNorth(x, y, rs)
	cgSouth := shared10.CanGoSouth(x, y, rs)
	cgEast := shared10.CanGoEast(x, y, rs)
	cgWest := shared10.CanGoWest(x, y, rs)

	if cgNorth && cgSouth {
		return '|'
	} else if cgNorth && cgEast {
		return 'L'
	} else if cgNorth && cgWest {
		return 'J'
	} else if cgSouth && cgEast {
		return 'F'
	} else if cgSouth && cgWest {
		return '7'
	} else if cgEast && cgWest {
		return '-'
	} else {
		panic("don't know what to replace S with")
	}
}

func GetGridWithoutUnusedPipes(sx, sy int, rs [][]rune) [][]rune {
	ret := make([][]rune, len(rs))
	for y := range rs {
		ret[y] = make([]rune, len(rs[0]))
	}

	const north = "north"
	const south = "south"
	const west = "west"
	const east = "east"

	x, y := sx, sy
	comeFrom := ""

	isFirst := true
	for {
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

		ret[y][x] = rs[y][x]

		if isFirst {
			isFirst = false
		} else if x == sx && y == sy {
			return ret
		}
	}
}

func CalculateNrOfEnclosedTiles(rs [][]rune) int {
	n := 0
	walls := 0
	var lastBend rune = 0

	// iterate horizontally over all tiles
	for y := range rs {
		for x := range rs[y] {
			r := rs[y][x]
			switch r {
			case '|':
				walls++
				continue
			case 'F', 'L':
				lastBend = r
				continue
			case '7':
				if lastBend == 'L' {
					walls++
				}
				continue
			case 'J':
				if lastBend == 'F' {
					walls++
				}
				continue
			case '-':
				continue

			default: // not a wall
				if walls%2 == 1 {
					rs[y][x] = 'e'
				}
			}
		}
	}

	// iterate vertically over all tiles
	for x := 0; x < len(rs[0]); x++ {
		for y := range rs {
			r := rs[y][x]
			switch r {
			case '-':
				walls++
				continue
			case 'F', '7':
				lastBend = r
				continue
			case 'L':
				if lastBend == '7' {
					walls++
				}
				continue
			case 'J':
				if lastBend == 'F' {
					walls++
				}
				continue
			case '|':
				continue

			case 'e': // only add enclosed tile when it is both vertically and horizontally enclosed
				if walls%2 == 1 {
					n++
				}
			}
		}
	}
	return n
}
