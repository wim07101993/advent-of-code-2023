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
	rs := shared.ReadAllRunesByLine(r)

	var sum = 0

	for y, row := range rs {
		for x, r := range row {
			if r != '*' {
				continue
			}
			ns := getAdjacentNumbers(rs, x, y)
			if len(ns) != 2 {
				continue
			}
			sum += ns[0] * ns[1]
		}
	}

	return sum
}

func getAdjacentNumbers(rs [][]rune, x int, y int) []int {
	canGoLeft := x > 0
	canGoRight := x+1 < len(rs[y])
	canGoUp := y > 0
	canGoDown := y+1 < len(rs)

	var ns []int

	checkRow := func(y int) {
		if unicode.IsDigit(rs[y][x]) {
			ns = append(ns, getNumber(rs[y], x))
		} else {
			if canGoLeft && unicode.IsDigit(rs[y][x-1]) {
				ns = append(ns, getNumber(rs[y], x-1))
			}
			if canGoRight && unicode.IsDigit(rs[y][x+1]) {
				ns = append(ns, getNumber(rs[y], x+1))
			}
		}
	}
	if canGoUp {
		checkRow(y - 1)
	}

	if canGoDown {
		checkRow(y + 1)
	}

	checkRow(y)

	return ns
}

func getNumber(runes []rune, x int) int {
	start := x
	for start > 0 && unicode.IsDigit(runes[start-1]) {
		start--
	}
	end := x
	for end < len(runes)-1 && unicode.IsDigit(runes[end+1]) {
		end++
	}
	i, err := strconv.Atoi(string(runes[start : end+1]))
	if err != nil {
		panic(err)
	}
	return i
}
