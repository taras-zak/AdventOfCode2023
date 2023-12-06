package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
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
	var totalSeeds int
	var seeds []rng
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
			seedRanges := strings.Split(strings.TrimSpace(value), " ")
			for i := 0; i < len(seedRanges)-1; i += 2 {
				startSeed, _ := strconv.Atoi(seedRanges[i])
				rangeLen, _ := strconv.Atoi(seedRanges[i+1])
				seeds = append(seeds, rng{start: startSeed, stop: startSeed + rangeLen})
				totalSeeds += rangeLen
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
	resChan := make(chan int)
	wg := sync.WaitGroup{}
	for _, r := range seeds {
		r := r
		wg.Add(1)
		go func() {
			defer wg.Done()
			var minLoc = int(^uint(0) >> 1)
			for seed := r.start; seed < r.stop; seed++ {
				loc := doMapping(seed, seedsToSoilRange)
				loc = doMapping(loc, soilToFertRange)
				loc = doMapping(loc, fertToWaterRange)
				loc = doMapping(loc, waterToLightRange)
				loc = doMapping(loc, lightToTempRange)
				loc = doMapping(loc, tempToHumRange)
				loc = doMapping(loc, humToLocRange)
				if loc < minLoc {
					minLoc = loc
				}
			}
			resChan <- minLoc
		}()
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()

	var minLoc = int(^uint(0) >> 1) // max int
	for localMin := range resChan {
		if localMin < minLoc {
			minLoc = localMin
		}
	}
	fmt.Println("Answer:", minLoc)
}

func doMapping(seed int, mapping []rng) int {
	left := 0
	right := len(mapping) - 1
	for left <= right {
		mid := left + (right-left)/2
		r := mapping[mid]
		if seed > r.stop {
			left = mid + 1
		} else if seed < r.start {
			right = mid - 1
		} else {
			return seed + r.diff
		}
	}
	return seed
	//for _, r := range mapping {
	//	if r.start <= seed && r.stop >= seed {
	//		seed = seed + r.diff
	//		break
	//	}
	//}
	//return seed
}

func parseRange(value string) []rng {
	var res []rng
	for _, rangeString := range strings.Split(strings.TrimSpace(value), "\n") {
		dst, _ := strconv.Atoi(strings.Split(rangeString, " ")[0])
		src, _ := strconv.Atoi(strings.Split(rangeString, " ")[1])
		diff, _ := strconv.Atoi(strings.Split(rangeString, " ")[2])
		res = append(res, rng{src, src + diff - 1, dst - src})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].start < res[j].start
	})
	return res
}
