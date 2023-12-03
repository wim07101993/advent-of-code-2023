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

	var power = 0

	for scanner.Scan() {
		g := shared.ParseGame(scanner.Text())
		red, green, blue := minCubesNeededForGame(g)
		power += red * green * blue

	}

	return power
}

func minCubesNeededForGame(g shared.Game) (red int, green int, blue int) {
	for _, r := range g.Rounds {
		if r.Red > red {
			red = r.Red
		}
		if r.Green > green {
			green = r.Green
		}
		if r.Blue > blue {
			blue = r.Blue
		}
	}
	return
}
