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
	for {
		var cubeQuantity = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		scanner.Scan()
		game := scanner.Text()
		if len(game) == 0 {
			break
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
					cubeQuantity[color] = count
				}
			}
		}
		localRes := 1
		for _, v := range cubeQuantity {
			localRes *= v
		}
		result += localRes
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your puzzle answer is:", result)
}
