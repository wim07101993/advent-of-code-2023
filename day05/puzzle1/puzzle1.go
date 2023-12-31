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

	var loc int
	for _, s := range seeds {
		if l := seedMap.GetSeedLocation(s); l < loc || loc == 0 {
			loc = l
		}
	}

	return loc
}

func ParseSeeds(s string) []int {
	ss := strings.Split(strings.Split(s, ": ")[1], " ")
	seeds := make([]int, len(ss))
	var err error
	for i := range ss {
		seeds[i], err = strconv.Atoi(ss[i])
		if err != nil {
			panic(err)
		}
	}
	return seeds
}
