package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rng struct {
	start, stop int
	diff        int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	res := strings.Split(string(input), "\n\n")
	var seeds []int
	var seedsToSoilRange []rng
	var soilToFertRange []rng
	var fertToWaterRange []rng
	var waterToLightRange []rng
	var lightToTempRange []rng
	var tempToHumRange []rng
	var humToLocRange []rng
	for _, part := range res {
		name := strings.Split(part, ":")[0]
		value := strings.Split(part, ":")[1]
		switch name {
		case "seeds":
			for _, seed := range strings.Split(value, " ") {
				if seed == "" {
					continue
				}
				s, _ := strconv.Atoi(seed)
				seeds = append(seeds, s)
			}
		case "seed-to-soil map":
			seedsToSoilRange = parseRange(value)
		case "soil-to-fertilizer map":
			soilToFertRange = parseRange(value)
		case "fertilizer-to-water map":
			fertToWaterRange = parseRange(value)
		case "water-to-light map":
			waterToLightRange = parseRange(value)
		case "light-to-temperature map":
			lightToTempRange = parseRange(value)
		case "temperature-to-humidity map":
			tempToHumRange = parseRange(value)
		case "humidity-to-location map":
			humToLocRange = parseRange(value)
		}
	}
	var minLoc = int(^uint(0) >> 1) // max int
	for _, seed := range seeds {
		seed = doMapping(seed, seedsToSoilRange)
		seed = doMapping(seed, soilToFertRange)
		seed = doMapping(seed, fertToWaterRange)
		seed = doMapping(seed, waterToLightRange)
		seed = doMapping(seed, lightToTempRange)
		seed = doMapping(seed, tempToHumRange)
		seed = doMapping(seed, humToLocRange)
		if seed < minLoc {
			minLoc = seed
		}
	}
	fmt.Println("Answer:", minLoc)
}

func doMapping(seed int, mapping []rng) int {
	for _, r := range mapping {
		if r.start <= seed && r.stop >= seed {
			seed = seed + r.diff
			break
		}
	}
	return seed
}

func parseRange(value string) []rng {
	var res []rng
	for _, rangeString := range strings.Split(strings.TrimSpace(value), "\n") {
		dst, _ := strconv.Atoi(strings.Split(rangeString, " ")[0])
		src, _ := strconv.Atoi(strings.Split(rangeString, " ")[1])
		diff, _ := strconv.Atoi(strings.Split(rangeString, " ")[2])
		res = append(res, rng{src, src + diff - 1, dst - src})
	}
	return res
}
