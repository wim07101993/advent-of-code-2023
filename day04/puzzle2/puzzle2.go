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

	cardCounts := map[int]int{}
	cardNr := 1
	for scanner.Scan() {
		cardCounts[cardNr]++
		n := GetNumberOfWinningNumbers(scanner.Text())
		for i := 1; i < n+1; i++ {
			cardCounts[cardNr+i] = cardCounts[cardNr+i] + cardCounts[cardNr]
		}
		cardNr++
	}

	var sum = 0
	for i, v := range cardCounts {
		if i >= cardNr {
			continue
		}
		sum += v
	}

	return sum
}
func GetNumberOfWinningNumbers(s string) int {
	ss := strings.Split(s, ": ")
	ss = strings.Split(ss[1], " | ")

	winning := shared.ParseNumbers(ss[0])
	card := shared.ParseNumbers(ss[1])

	c := 0
	for _, n := range card {
		if !slices.Contains(winning, n) {
			continue
		}
		c++
	}
	return c
}
