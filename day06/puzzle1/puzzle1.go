package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
		result *= Magic(ts[i], ds[i])
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

// Magic does some magic algorithm which is explained in maths.md
func Magic(t int, s int) int {
	ft := float64(t)
	fs := float64(s)
	trecord1 := (-ft + math.Sqrt(ft*ft-4*fs)) / (-2)
	trecord2 := (-ft - math.Sqrt(ft*ft-4*fs)) / (-2)
	n := math.Ceil(trecord2-1) - math.Floor(trecord1+1) + 1
	return int(n)
}
