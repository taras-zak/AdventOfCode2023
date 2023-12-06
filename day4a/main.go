package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("show me your puzzle:")
	scanner := bufio.NewScanner(os.Stdin)
	var result int
	for {
		scanner.Scan()
		game := scanner.Text()
		if len(game) == 0 {
			break
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
		gameRes := 1
		for _, num := range strings.Split(myNumbersString, " ") {
			if _, ok := winningNumbersMap[num]; ok {
				gameRes *= 2
			}
		}
		if gameRes != 1 {
			result += gameRes / 2
		}
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your puzzle answer is:", result)
}
