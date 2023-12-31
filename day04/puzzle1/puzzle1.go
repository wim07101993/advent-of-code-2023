package main

import (
	"advent-of-code-2023/day04/shared"
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var sum = 0
	for scanner.Scan() {
		sum += SolveLine(scanner.Text())
	}

	return sum
}

func SolveLine(s string) int {
	ss := strings.Split(s, ": ")
	ss = strings.Split(ss[1], " | ")

	winning := shared.ParseNumbers(ss[0])
	card := shared.ParseNumbers(ss[1])

	pts := 0
	for _, n := range card {
		if !slices.Contains(winning, n) {
			continue
		}
		if pts == 0 {
			pts = 1
		} else {
			pts *= 2
		}
	}
	return pts
}
