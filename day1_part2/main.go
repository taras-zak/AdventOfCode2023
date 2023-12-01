package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("show me your puzzle:")
	scanner := bufio.NewScanner(os.Stdin)
	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	reverseReplacer := strings.NewReplacer(
		"oneight", "8",
		"twone", "1",
		"threeight", "8",
		"fiveight", "8",
		"sevenine", "9",
		"eightwo", "2",
		"eighthree", "3",
		"nineight", "8",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	var result int
	for {
		scanner.Scan()
		origLine := scanner.Text()
		if len(origLine) == 0 {
			break
		}

		line := []rune(replacer.Replace(origLine))
		localRes := 0
		for _, r := range line {
			if unicode.IsDigit(r) {
				localRes = int(r-'0') * 10
				break
			}
		}
		line = []rune(reverseReplacer.Replace(origLine))
		for i := len(line) - 1; i >= 0; i-- {
			r := line[i]
			if unicode.IsDigit(r) {
				localRes += int(r - '0')
				break
			}
		}
		result += localRes
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your puzzle answer is: ", result)
}
