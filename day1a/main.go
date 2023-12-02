package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	fmt.Println("show me your puzzle:")
	scanner := bufio.NewScanner(os.Stdin)

	var result int
	for {
		scanner.Scan()
		line := []rune(scanner.Text())
		if len(line) == 0 {
			break
		}
		localRes := 0
		for _, r := range line {
			if unicode.IsDigit(r) {
				localRes = int(r-'0') * 10
				break
			}
		}
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
