package main

import (
	"bufio"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"regexp"
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

	var cs []chan string
	for scanner.Scan() {
		c := make(chan string)
		cs = append(cs, c)
		fmt.Println("start searching for", scanner.Text())
		go func(l string, c chan<- string) {
			ss := strings.Split(l, " ")
			damaged := UnfoldDamaged(ParseDamaged(strings.Split(ss[1], ",")))
			unfolded := UnfoldSprings(ss[0])
			GetArrangements(unfolded, damaged, c)
		}(scanner.Text(), c)
	}

	pb := progressbar.Default(int64(len(cs)))

	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	var total int64
	for _, c := range cs {
		go func(c <-chan string) {
			sum := 0
			for range c {
				sum++
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

func GetArrangements(springs string, damaged []int, results chan<- string) {
	switch len(damaged) {
	case 0:
		// if no damaged springs are left, solution is to replace all '?' with '.'
		results <- strings.Replace(springs, "?", ".", -1)
		close(results)
		return
	case 1:
		GetLastArrangements(springs, damaged[0], results)
		return
	}

	defer func() {
		close(results)
	}()

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

		c := make(chan string)

		go func() {
			GetArrangements(springs[end+1:], damaged[1:], c)
		}()

		for arr := range c {
			arr = springs[:start] + strings.Repeat("#", damaged[0]) + "." + arr
			arr = strings.Replace(arr, "?", ".", -1)
			results <- arr
		}
	}
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

func GetLastArrangements(springs string, damaged int, results chan<- string) {
	defer func() {
		close(results)
	}()

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

		results <- arr
	}
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
