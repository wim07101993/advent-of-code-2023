package main

import (
	"advent-of-code-2023/day02/shared"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	const red = 12
	const green = 13
	const blue = 14

	var sum = 0

	for scanner.Scan() {
		g := shared.ParseGame(scanner.Text())
		if isGamePossibleWith(g, red, green, blue) {
			sum += g.Id
		}
	}

	return sum
}

func isGamePossibleWith(g shared.Game, red int, green int, blue int) bool {
	for _, r := range g.Rounds {
		if r.Red > red {
			return false
		}
		if r.Green > green {
			return false
		}
		if r.Blue > blue {
			return false
		}
	}
	return true
}
