package shared

import (
	"strconv"
	"strings"
)

func ParseNumbers(s string) []int {
	ss := strings.Split(s, " ")
	ns := make([]int, 0, len(ss))
	for _, n := range ss {
		if n == "" {
			continue
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		ns = append(ns, i)
	}
	return ns
}
