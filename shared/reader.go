package shared

import (
	"bufio"
	"io"
)

func ReadAllRunesByLine(r io.Reader) [][]rune {
	var runes [][]rune

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		runes = append(runes, []rune(scanner.Text()))
	}

	return runes
}
