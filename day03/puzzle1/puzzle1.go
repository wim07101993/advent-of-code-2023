package main

import (
	"advent-of-code-2023/shared"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	runes := shared.ReadAllRunesByLine(r)

	var sum = 0

	for y, row := range runes {
		var foundDigit bool
		for x, r := range row {
			if !unicode.IsDigit(r) {
				foundDigit = false
				continue
			}

			if !foundDigit {
				n := getNumber(row, x)
				if checkNumberHasAdjacentSymbols(runes, n, x, y) {
					i, err := strconv.Atoi(n)
					if err != nil {
						panic(err)
					}
					sum += i
				}
			}

			foundDigit = true
		}
	}

	return sum
}

func checkNumberHasAdjacentSymbols(rss [][]rune, n string, x int, y int) bool {
	var startX int
	var endX int

	if x > 0 {
		startX = x - 1
	} else {
		startX = x
	}
	endRow := len(rss[y]) - 1
	endNumber := x + len(n) - 1
	if endNumber < endRow {
		endX = endNumber + 1
	} else {
		endX = endNumber
	}

	if y > 0 {
		// iterate over the row above number
		for x := startX; x < endX+1; x++ {
			r := rss[y-1][x]
			if !unicode.IsDigit(r) && r != '.' {
				return true
			}
		}
	}

	if y < len(rss)-1 {
		// iterate over the row below number
		for x := startX; x < endX+1; x++ {
			r := rss[y+1][x]
			if !unicode.IsDigit(r) && r != '.' {
				return true
			}
		}
	}

	if startX != x {
		// check rune before number
		r := rss[y][startX]
		if !unicode.IsDigit(r) && r != '.' {
			return true
		}
	}
	if endX < endRow {
		// check rune after number
		r := rss[y][endX]
		if !unicode.IsDigit(r) && r != '.' {
			return true
		}
	}

	return false
}

func getNumber(runes []rune, start int) string {
	for x := start; x < len(runes); x++ {
		if !unicode.IsDigit(runes[x]) {
			return string(runes[start:x])
		}
	}
	return string(runes[start:])
}
