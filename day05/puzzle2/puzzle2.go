package main

import (
	"advent-of-code-2023/day05/shared"
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

type SeedRange struct {
	start  int
	length int
}

func Solve(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	seeds := ParseSeeds(scanner.Text())
	seedMap := shared.SeedMap{}

	for scanner.Scan() {
		name := scanner.Text()
		if name == "" {
			continue
		}
		ms := shared.ParseMappers(scanner)
		switch name {
		case "seed-to-soil map:":
			seedMap.SeedToSoil = ms
		case "soil-to-fertilizer map:":
			seedMap.SoilToFertilizer = ms
		case "fertilizer-to-water map:":
			seedMap.FertilizerToWater = ms
		case "water-to-light map:":
			seedMap.WaterToLight = ms
		case "light-to-temperature map:":
			seedMap.LightToTemperature = ms
		case "temperature-to-humidity map:":
			seedMap.TemperatureToHumidity = ms
		case "humidity-to-location map:":
			seedMap.HumidityToLocation = ms
		}
	}

	cs := make([]chan int, len(seeds))
	for i, sr := range seeds {
		cs[i] = make(chan int, 10)
		go sr.GetLocations(seedMap, cs[i])
	}

	return chanMin(cs)
}

func ParseSeeds(s string) []SeedRange {
	ss := strings.Split(strings.Split(s, ": ")[1], " ")
	seeds := make([]SeedRange, len(ss)/2)
	for i := range seeds {
		start, err := strconv.Atoi(ss[i*2])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(ss[(i*2)+1])
		if err != nil {
			panic(err)
		}
		seeds[i] = SeedRange{start, length}
	}
	return seeds
}

func (sr SeedRange) GetLocations(m shared.SeedMap, c chan<- int) {
	l := 0
	for s := sr.start; s < sr.start+sr.length; s++ {
		if x := m.GetSeedLocation(s); x < l || l == 0 {
			l = x
		}
	}
	c <- l
	close(c)
}

func chanMin(cs []chan int) int {
	x := 0
	for _, c := range cs {
		for y := range c {
			if y < x || x == 0 {
				x = y
			}
		}
	}
	return x
}
