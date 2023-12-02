package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digits = map[string]rune{
	"zero":  '0',
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

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

func SolveLine(l string) int {
	var dig1 rune
	for i := range l {
		var isDigit bool
		dig1, isDigit = checkRune(l, i)
		if isDigit {
			break
		}
	}

	if dig1 == 0 {
		panic("no digits found")
	}

	var dig2 rune
	for i := len(l) - 1; i >= 0; i-- {
		var isDigit bool
		dig2, isDigit = checkRune(l, i)
		if isDigit {
			break
		}
	}

	num, err := strconv.Atoi(string([]rune{dig1, dig2}))
	if err != nil {
		panic(err)
	}
	return num
}

func checkRune(l string, i int) (rune, bool) {
	if unicode.IsNumber(rune(l[i])) {
		return rune(l[i]), true
	}
	s := l[i:]
	for name, value := range digits {
		if strings.HasPrefix(s, name) {
			return value, true
		}
	}
	return 0, false
}
