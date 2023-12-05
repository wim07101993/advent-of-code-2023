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

type SeedMap struct {
	SeedToSoil            Mappers
	SoilToFertilizer      Mappers
	FertilizerToWater     Mappers
	WaterToLight          Mappers
	LightToTemperature    Mappers
	TemperatureToHumidity Mappers
	HumidityToLocation    Mappers
}
type Mapper struct {
	StartSource      int
	StartDestination int
	RangeLength      int
}
type Mappers []Mapper

func Solve(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	seeds := ParseSeeds(scanner.Text())
	seedMap := SeedMap{}

	for scanner.Scan() {
		name := scanner.Text()
		if name == "" {
			continue
		}
		ms := ParseMappers(scanner)
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

func ParseMappers(scanner *bufio.Scanner) Mappers {
	var ms []Mapper
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			return ms
		}

		ss := strings.Split(l, " ")
		dest, err := strconv.Atoi(ss[0])
		if err != nil {
			panic(err)
		}

		source, err := strconv.Atoi(ss[1])
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(ss[2])
		if err != nil {
			panic(err)
		}

		ms = append(ms, Mapper{
			StartSource:      source,
			StartDestination: dest,
			RangeLength:      r,
		})
	}
	return ms
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

func (sm SeedMap) GetSeedLocation(seed int) int {
	soil := sm.SeedToSoil.GetDestination(seed)
	fert := sm.SoilToFertilizer.GetDestination(soil)
	water := sm.FertilizerToWater.GetDestination(fert)
	light := sm.WaterToLight.GetDestination(water)
	temp := sm.LightToTemperature.GetDestination(light)
	hum := sm.TemperatureToHumidity.GetDestination(temp)
	return sm.HumidityToLocation.GetDestination(hum)
}

func (ms Mappers) GetDestination(source int) int {
	for _, m := range ms {
		if source >= m.StartSource && source < m.StartSource+m.RangeLength {
			x := source - m.StartSource
			return m.StartDestination + x
		}
	}
	return source
}
