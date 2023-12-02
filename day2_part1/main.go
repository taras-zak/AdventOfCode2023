package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cubeQuantity = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	fmt.Println("show me your puzzle:")
	scanner := bufio.NewScanner(os.Stdin)
	var result int
gameLoop:
	for {
		scanner.Scan()
		game := scanner.Text()
		if len(game) == 0 {
			break
		}
		gameString := strings.Split(game, ":")[0]
		gameIDString := strings.Split(gameString, " ")[1]
		gameID, err := strconv.Atoi(gameIDString)
		if err != nil {
			panic(err)
		}
		setsString := strings.Split(game, ":")[1]
		sets := strings.Split(setsString, ";")
		for _, setString := range sets {
			cubes := strings.Split(setString, ",")
			for _, cubeString := range cubes {
				countString := strings.Split(strings.TrimSpace(cubeString), " ")[0]
				count, err := strconv.Atoi(countString)
				if err != nil {
					panic(err)
				}
				color := strings.Split(strings.TrimSpace(cubeString), " ")[1]
				if q, ok := cubeQuantity[color]; !ok || count > q {
					continue gameLoop
				}
			}
		}
		result += gameID
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your puzzle answer is:", result)
}
