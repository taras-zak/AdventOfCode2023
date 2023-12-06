package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type coord struct {
	row, col int
}

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
	puzzleGears := make(map[coord][]int)

	rows := len(puzzle)
	cols := len(puzzle[0])
	for row := 0; row < rows; row++ {
		adjacentGears := make(map[coord]bool)
		var number []rune
		for col := 0; col < cols; col++ {
			if unicode.IsDigit(puzzle[row][col]) {
				for _, c := range checkGears(puzzle, row, col) {
					adjacentGears[c] = true
				}
				number = append(number, puzzle[row][col])
				// check if end of number
				if col+1 > cols-1 || !unicode.IsDigit(puzzle[row][col+1]) {
					n, _ := strconv.Atoi(string(number))
					for gear, _ := range adjacentGears {
						puzzleGears[gear] = append(puzzleGears[gear], n)
					}
					number = nil
					adjacentGears = make(map[coord]bool)
				}
			}
		}
	}
	for _, numbers := range puzzleGears {
		if len(numbers) == 2 {
			res += numbers[0] * numbers[1]
		}
	}
	fmt.Println("Your puzzle answer is:", res)
}

func checkGears(puzzle [][]rune, row, col int) []coord {
	res := make([]coord, 0)
	// left
	if col > 1 && isGear(puzzle[row][col-1]) {
		res = append(res, coord{row, col - 1})
	}
	// left up
	if col > 1 && row > 1 && isGear(puzzle[row-1][col-1]) {
		res = append(res, coord{row - 1, col - 1})
	}
	// up
	if row > 1 && isGear(puzzle[row-1][col]) {
		res = append(res, coord{row - 1, col})
	}
	// right up
	if row > 1 && col < len(puzzle[0])-1 && isGear(puzzle[row-1][col+1]) {
		res = append(res, coord{row - 1, col + 1})
	}
	// right
	if col < len(puzzle[0])-1 && isGear(puzzle[row][col+1]) {
		res = append(res, coord{row, col + 1})
	}
	// right down
	if col < len(puzzle[0])-1 && row < len(puzzle)-1 && isGear(puzzle[row+1][col+1]) {
		res = append(res, coord{row + 1, col + 1})
	}
	// down
	if row < len(puzzle)-1 && isGear(puzzle[row+1][col]) {
		res = append(res, coord{row + 1, col})
	}
	// left down
	if col > 1 && row < len(puzzle)-1 && isGear(puzzle[row+1][col-1]) {
		res = append(res, coord{row + 1, col - 1})
	}
	return res
}

func isGear(r rune) bool {
	if r == '*' {
		return true
	}
	return false
}
