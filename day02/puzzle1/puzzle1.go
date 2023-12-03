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

	const red = 12
	const green = 13
	const blue = 14

	var sum = 0

	for scanner.Scan() {
		g := ParseGame(scanner.Text())
		if g.isPossibleWith(red, green, blue) {
			sum += g.id
		}
	}

	return sum
}

func ParseGame(s string) Game {
	ss := strings.Split(s, ": ")
	id, err := strconv.Atoi(strings.Split(ss[0], " ")[1])
	if err != nil {
		panic(err)
	}
	g := Game{id: id}
	rounds := strings.Split(ss[1], "; ")
	g.rounds = make([]Round, len(rounds), len(rounds))
	for i, round := range rounds {
		g.rounds[i] = parseRound(round)
	}
	return g
}

func parseRound(s string) Round {
	ss := strings.Split(s, ", ")
	round := Round{}
	for _, die := range ss {
		name, count := parseDie(die)
		switch name {
		case "red":
			round.red = count
		case "green":
			round.green = count
		case "blue":
			round.blue = count
		}
	}
	return round
}

func parseDie(s string) (name string, count int) {
	ss := strings.Split(s, " ")
	count, err := strconv.Atoi(ss[0])
	if err != nil {
		panic(err)
	}
	return ss[1], count
}

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func (g *Game) isPossibleWith(red int, green int, blue int) bool {
	for _, r := range g.rounds {
		if r.red > red {
			return false
		}
		if r.green > green {
			return false
		}
		if r.blue > blue {
			return false
		}
	}
	return true
}
