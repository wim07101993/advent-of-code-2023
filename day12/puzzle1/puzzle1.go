package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
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
		ss := strings.Split(scanner.Text(), " ")
		damaged := ParseAll(strings.Split(ss[1], ","))
		sum += len(GetArrangements(ss[0], damaged))
	}

	return sum
}

func GetArrangements(springs string, damaged []int) []string {
	switch len(damaged) {
	case 0:
		// if no damaged springs are left, solution is to replace all '?' with '.'
		return []string{strings.Replace(springs, "?", ".", -1)}
	case 1:
		return GetLastArrangements(springs, damaged[0])
	}

	var arrs []string

	r := GetArrangementRegex(damaged)
	for start := 0; start < len(springs)-damaged[0]-1; start++ {
		end := start + damaged[0]
		sub := springs[start:end]
		trailing := springs[end]

		// once we are past the first #, no more possible arrangements
		if start > 0 && springs[start-1] == '#' {
			break
		}

		// sub-string should only contain '#' and '?' and be followed by a '.'
		if (trailing != '.' && trailing != '?') || strings.Contains(sub, ".") {
			continue
		}

		for _, arr := range GetArrangements(springs[end+1:], damaged[1:]) {
			arr = springs[:start] + strings.Repeat("#", damaged[0]) + "." + arr
			arr = strings.Replace(arr, "?", ".", -1)

			if !slices.Contains(arrs, arr) && r.MatchString(arr) {
				arrs = append(arrs, arr)
			}
		}
	}

	return arrs
}

func GetArrangementRegex(damaged []int) *regexp.Regexp {
	b := strings.Builder{}
	b.WriteString("\\.*")
	for i := 0; i < len(damaged)-1; i++ {
		b.WriteString(fmt.Sprintf("#{%v}\\.+", damaged[i]))
	}
	b.WriteString(fmt.Sprintf("#{%v}", damaged[len(damaged)-1]))
	return regexp.MustCompile(b.String())
}

func GetLastArrangements(springs string, damaged int) []string {
	var arrs []string
	for start := 0; start < len(springs)-damaged+1; start++ {
		end := start + damaged
		sub := springs[start:end]

		// once we are past the first #, no more possible arrangements
		if start > 0 && springs[start-1] == '#' {
			break
		}

		// sub-string should only contain '#' and '?'
		if strings.Contains(sub, ".") {
			continue
		}

		// part after the substring should not contain any '#' anymore
		trailing := springs[end:]
		if strings.Contains(trailing, "#") {
			continue
		}

		// replace the substring with # and add it to the arrs
		arr := springs[:start] + strings.Repeat("#", damaged) + trailing
		arr = strings.Replace(arr, "?", ".", -1)
		arrs = append(arrs, arr)
	}

	return arrs
}

func ParseAll(ss []string) []int {
	is := make([]int, len(ss))
	var err error
	for i, s := range ss {
		is[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return is
}
