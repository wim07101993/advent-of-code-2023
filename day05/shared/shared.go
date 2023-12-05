package shared

import (
	"bufio"
	"strconv"
	"strings"
)

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
