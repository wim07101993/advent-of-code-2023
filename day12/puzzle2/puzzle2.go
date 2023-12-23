package main

import (
	"bufio"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var cs []chan int64
	for scanner.Scan() {
		c := make(chan int64)
		cs = append(cs, c)
		fmt.Println("start searching for", scanner.Text())
		go func(l string, c chan<- int64) {
			ss := strings.Split(l, " ")
			damaged := UnfoldDamaged(ParseDamaged(strings.Split(ss[1], ",")))
			unfolded := UnfoldSprings(ss[0])
			GetArrangementCounts(unfolded, damaged, c)
		}(scanner.Text(), c)
	}

	pb := progressbar.Default(int64(len(cs)))

	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	var total int64
	for _, c := range cs {
		go func(c <-chan int64) {
			var sum int64
			for count := range c {
				sum += count
			}
			total += int64(sum)
			pb.Describe(fmt.Sprintf("%v goroutines\t%v total", runtime.NumGoroutine(), total))
			err := pb.Add(1)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(c)
	}

	wg.Wait()
	return total
}

func GetArrangementCounts(springs string, damaged []int, results chan<- int64) {
	defer func() {
		close(results)
	}()
	switch len(damaged) {
	case 0:
		// if no damaged springs are left, solution is to replace all '?' with '.'
		results <- 1
		return
	case 1:
		results <- int64(GetLastArrangementsCount(springs, damaged[0]))
		return
	}

	for i := 0; i < len(springs)-damaged[0]-1; i++ {
		start := i
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

		c := make(chan int64)

		go func() {
			GetArrangementCounts(springs[end+1:], damaged[1:], c)
		}()

		for count := range c {
			results <- count
		}
	}
}

func GetLastArrangementsCount(springs string, damaged int) int {
	count := 0
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

		count++
	}
	return count
}

func ParseDamaged(ss []string) []int {
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

func UnfoldSprings(s string) string {
	b := strings.Builder{}
	for i := 0; i < 4; i++ {
		b.WriteString(s)
		b.WriteRune('?')
	}
	b.WriteString(s)
	return b.String()
}

func UnfoldDamaged(ds []int) []int {
	ret := make([]int, len(ds)*5)
	for i := range ret {
		ret[i] = ds[i%len(ds)]
	}
	return ret
}
