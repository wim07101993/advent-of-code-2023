package main

import (
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

	sum := 0

	for scanner.Scan() {
		l := scanner.Text()
		hist := ParseHistory(l)
		sum += Extrapolate(hist)
	}

	return sum
}

func ParseHistory(s string) []int {
	ss := strings.Split(s, " ")
	hist := make([]int, len(ss))
	for i := range ss {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			panic(err)
		}
		hist[i] = n
	}
	return hist
}

func Extrapolate(ns []int) int {
	difs := make([]int, len(ns)-1)
	allZero := true
	for i := 0; i < len(ns)-1; i++ {
		difs[i] = ns[i+1] - ns[i]
		if difs[i] != 0 {
			allZero = false
		}
	}
	if !allZero {
		nextDif := Extrapolate(difs)
		return ns[0] - nextDif
	} else {
		return ns[0]
	}
}
