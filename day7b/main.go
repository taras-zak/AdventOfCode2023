package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards []int
	rank  int
	bid   int
}

var CardRanks = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func main() {
	fmt.Println("show me your input:")
	scanner := bufio.NewScanner(os.Stdin)
	var result int
	var hands []hand
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		cards := strings.Split(line, " ")[0]
		bid, _ := strconv.Atoi(strings.Split(line, " ")[1])
		countMap := make(map[rune]int)

		var ranks []int
		for _, card := range cards {
			rank := CardRanks[card]
			ranks = append(ranks, rank)
			countMap[card]++
		}
		var counts []int
		var jokerCounts int
		for k, v := range countMap {
			if k == 'J' {
				jokerCounts = v
			}
			counts = append(counts, v)
		}
		sort.Ints(counts)

		rank := 0
		switch len(countMap) {
		case 1:
			rank = 50
		case 2:
			if counts[len(counts)-1] == 4 {
				rank = 40
			}
			if counts[len(counts)-1] == 3 && counts[len(counts)-2] == 2 {
				rank = 35
			}
			if jokerCounts > 0 {
				rank = 50
			}
		case 3:
			if counts[len(counts)-1] == 3 {
				rank = 30
				if jokerCounts > 0 {
					rank = 40
				}
			}
			if counts[len(counts)-1] == 2 && counts[len(counts)-2] == 2 {
				rank = 25
				if jokerCounts == 2 {
					rank = 40
				}
				if jokerCounts == 1 {
					rank = 35
				}
			}
		case 4:
			rank = 20
			if jokerCounts > 0 {
				rank = 30
			}
		case 5:
			rank = 19
			if jokerCounts > 0 {
				rank = 20
			}
		}
		hands = append(hands, hand{ranks, rank, bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank == hands[j].rank {
			for cardIdx := 0; cardIdx <= 5; cardIdx++ {
				if hands[i].cards[cardIdx] != hands[j].cards[cardIdx] {
					return hands[i].cards[cardIdx] < hands[j].cards[cardIdx]
				}
			}
		}
		return hands[i].rank < hands[j].rank
	})
	for i, h := range hands {
		result += h.bid * (i + 1)
		fmt.Println("hand:", h)

	}
	fmt.Println("Your answer is:", result)

}
