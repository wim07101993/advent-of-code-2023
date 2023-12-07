package main

import (
	"advent-of-code-2023/day06/shared"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	ts := ParseNumbers(strings.Split(scanner.Text(), ":")[1])
	scanner.Scan()
	ds := ParseNumbers(strings.Split(scanner.Text(), ":")[1])

	if len(ts) != len(ds) {
		panic("Expected as much times as distances...")
	}

	result := 1

	for i := range ts {
		result *= shared.Magic(ts[i], ds[i])
	}

	return result
}

func ParseNumbers(s string) []int {
	ss := strings.Split(s, " ")
	var is []int
	for _, si := range ss {
		if si == "" {
			continue
		}
		i, err := strconv.Atoi(si)
		if err != nil {
			panic(err)
		}
		is = append(is, i)
	}
	return is
}
