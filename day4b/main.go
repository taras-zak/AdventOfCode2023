package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("show me your puzzle:")
	scanner := bufio.NewScanner(os.Stdin)
	var result int
	copies := make(map[int]int)
	for {
		scanner.Scan()
		game := scanner.Text()
		if len(game) == 0 {
			break
		}
		gameString := strings.Split(game, ":")[0]
		cardIDString := strings.Split(gameString, "Card")[1]
		cardID, err := strconv.Atoi(strings.TrimSpace(cardIDString))
		if err != nil {
			panic(err)
		}
		cardString := strings.Split(game, ": ")[1]
		winningNumbersString := strings.Split(cardString, "| ")[0]
		winningNumbersMap := make(map[string]bool)
		for _, num := range strings.Split(winningNumbersString, " ") {
			if num == "" {
				continue
			}
			winningNumbersMap[num] = true
		}
		myNumbersString := strings.Split(cardString, "| ")[1]
		copies[cardID]++
		idToCopy := cardID + 1
		for _, num := range strings.Split(myNumbersString, " ") {
			if _, ok := winningNumbersMap[num]; ok {
				copies[idToCopy] += 1 * copies[cardID]
				idToCopy++
			}
		}

	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	for _, count := range copies {
		result += count
	}
	fmt.Println("Your puzzle answer is:", result)
}
