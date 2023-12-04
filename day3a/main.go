package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("show me your puzzle:")
	scanner := bufio.NewScanner(os.Stdin)
	var puzzle [][]rune
	var res int
	for {
		scanner.Scan()
		row := scanner.Text()
		if len(row) == 0 {
			break
		}
		puzzle = append(puzzle, []rune(row))
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	var testNumbers []string
	rows := len(puzzle)
	cols := len(puzzle[0])
	for row := 0; row < rows; row++ {
		hasSymbol := false
		var number []rune
		for col := 0; col < cols; col++ {
			hasSymbol = hasSymbol || checkHasSymbol(puzzle, row, col)
			if unicode.IsDigit(puzzle[row][col]) {
				number = append(number, puzzle[row][col])
				// check if end of number
				if col+1 > cols-1 || !unicode.IsDigit(puzzle[row][col+1]) {
					isNextColHasSymbol := false
					if col+1 <= cols-1 {
						isNextColHasSymbol = checkHasSymbol(puzzle, row, col+1)
					}
					if hasSymbol || isNextColHasSymbol {
						testNumbers = append(testNumbers, string(number))
						n, _ := strconv.Atoi(string(number))
						res += n
					}
					hasSymbol = false
					number = nil
				}
			}
			if col < cols-1 && !unicode.IsDigit(puzzle[row][col+1]) {
				hasSymbol = false
			}
		}
	}
	fmt.Println("Your puzzle answer is:", res)
}

func checkHasSymbol(puzzle [][]rune, row, col int) bool {
	// up
	if row > 1 && isSymbol(puzzle[row-1][col]) {
		return true
	}
	// curr
	if isSymbol(puzzle[row][col]) {
		return true
	}
	// down
	if row < len(puzzle)-1 && isSymbol(puzzle[row+1][col]) {
		return true
	}
	return false
}

func isSymbol(r rune) bool {
	if r == '.' {
		return false
	}
	if unicode.IsLetter(r) {
		return false
	}
	if unicode.IsNumber(r) {
		return false
	}
	return true
}
