package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var sum = 0
	for scanner.Scan() {
		x := SolveLine(scanner.Text())
		fmt.Printf("%v + %v = %v\n", sum, x, sum+x)
		sum += SolveLine(scanner.Text())
	}

	return sum
}
func SolveLine(l string) int {
	var dig1 rune
	for _, r := range l {
		if unicode.IsNumber(r) {
			dig1 = r
			break
		}
	}
	var dig2 rune
	for i := len(l) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(l[i])) {
			dig2 = rune(l[i])
			break
		}
	}

	num, err := strconv.Atoi(string([]rune{dig1, dig2}))
	if err != nil {
		panic(err)
	}
	return num
}
