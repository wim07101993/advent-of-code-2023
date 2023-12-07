package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hands []Hand

type Hand struct {
	Strength string
	Bid      int
	Type     HandType
}

type HandType int
type CardLabel rune

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var hands Hands
	for scanner.Scan() {
		hands = append(hands, ParseHand(scanner.Text()))
	}

	hands.Sort()

	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.Bid
	}
	return sum
}

func ParseHand(s string) Hand {
	ss := strings.Split(s, " ")
	bid, err := strconv.Atoi(ss[1])
	if err != nil {
		panic(err)
	}
	return Hand{ss[0], bid, ParseHandType(ss[0])}
}

func ParseHandType(s string) HandType {
	m := map[rune]int{}
	jokerCount := 0
	for _, r := range s {
		if r == 'J' {
			jokerCount++
		}
		m[r]++
	}
	if jokerCount == 5 || jokerCount == 4 {
		return FiveOfAKind
	}

	hasThree := false
	hasPair := false
	for k, v := range m {
		if k == 'J' {
			continue
		}
		switch v {
		case 5:
			return FiveOfAKind
		case 4:
			if jokerCount > 0 {
				return FiveOfAKind
			}
			return FourOfAKind
		case 3:
			if jokerCount == 2 {
				return FiveOfAKind
			}
			if jokerCount == 1 {
				return FourOfAKind
			}
			if hasPair {
				return FullHouse
			}
			hasThree = true
		case 2:
			if jokerCount == 3 {
				return FiveOfAKind
			}
			if jokerCount == 2 {
				return FourOfAKind
			}
			if (jokerCount == 1 && hasPair) || hasThree {
				return FullHouse
			}
			if hasPair {
				return TwoPair
			}
			hasPair = true
		}
	}

	if hasThree {
		return ThreeOfAKind
	}
	if hasPair {
		if jokerCount == 1 {
			return ThreeOfAKind
		}
		return OnePair
	}

	switch jokerCount {
	case 3:
		return FourOfAKind
	case 2:
		return ThreeOfAKind
	case 1:
		return OnePair
	default:
		return HighCard
	}
}

func (l CardLabel) Value() int {
	switch l {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'T':
		return 10
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	case 'J':
		return 1
	}
	panic(fmt.Errorf("unknown card label '%v'", l))
}

func (h1 Hand) IsWeakerThan(h2 Hand) bool {
	if h1.Type < h2.Type {
		return true
	}
	if h1.Type > h2.Type {
		return false
	}

	for i := range h1.Strength {
		v1 := CardLabel(h1.Strength[i]).Value()
		v2 := CardLabel(h2.Strength[i]).Value()
		if v1 < v2 {
			return true
		}
		if v1 > v2 {
			return false
		}
	}

	return false
}

func (hs Hands) Sort() {
	sort.Slice(hs, func(i, j int) bool {
		return hs[i].IsWeakerThan(hs[j])
	})
}
