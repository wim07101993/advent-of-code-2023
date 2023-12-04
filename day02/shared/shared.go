package shared

import (
	"strconv"
	"strings"
)

func ParseGame(s string) Game {
	ss := strings.Split(s, ": ")
	id, err := strconv.Atoi(strings.Split(ss[0], " ")[1])
	if err != nil {
		panic(err)
	}
	g := Game{Id: id}
	rounds := strings.Split(ss[1], "; ")
	g.Rounds = make([]Round, len(rounds))
	for i, round := range rounds {
		g.Rounds[i] = parseRound(round)
	}
	return g
}

func parseRound(s string) Round {
	ss := strings.Split(s, ", ")
	round := Round{}
	for _, die := range ss {
		color, count := parseCubeCount(die)
		switch color {
		case "red":
			round.Red = count
		case "green":
			round.Green = count
		case "blue":
			round.Blue = count
		}
	}
	return round
}

func parseCubeCount(s string) (name string, count int) {
	ss := strings.Split(s, " ")
	count, err := strconv.Atoi(ss[0])
	if err != nil {
		panic(err)
	}
	return ss[1], count
}

type Game struct {
	Id     int
	Rounds []Round
}

type Round struct {
	Red   int
	Green int
	Blue  int
}
