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
	t := ParseNumber(strings.Split(scanner.Text(), ":")[1])
	scanner.Scan()
	d := ParseNumber(strings.Split(scanner.Text(), ":")[1])

	return shared.Magic(t, d)
}

func ParseNumber(s string) int {
	si := strings.Replace(s, " ", "", -1)
	i, err := strconv.Atoi(si)
	if err != nil {
		panic(err)
	}
	return i
}
