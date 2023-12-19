package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Show me your puzzle:")
	scanner := bufio.NewScanner(os.Stdin)

	var res int
	var matrix [][]rune
	for {
		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		if line == "" {
			break
		}
		matrix = append(matrix, []rune(line))
	}
	rotateMatrix(matrix, false)
	//idx := make(map[int][][]rune)
	period := 0
	for cycle := 0; cycle < 1_000_000_000; cycle++ {
		matrix = slideMatrix(matrix)
		rotateMatrix(matrix, true)
		matrix = slideMatrix(matrix)
		rotateMatrix(matrix, true)
		matrix = slideMatrix(matrix)
		rotateMatrix(matrix, true)
		matrix = slideMatrix(matrix)
		rotateMatrix(matrix, true)

		after := matrixToString(matrix)
		if start, ok := cache[after]; ok {
			period = cycle - start
			cycle += (1_000_000_000 - cycle) / period * period
		} else {
			cache[after] = cycle
		}

	}
	res = calcRes(matrix)
	fmt.Println("Answer:", res)
}

func calcRes(matrix [][]rune) int {
	var res int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 'O' {
				res += len(matrix[row]) - col
			}
		}

	}
	return res
}

var cache = make(map[string]int)

func slideMatrix(matrix [][]rune) [][]rune {
	for i := 0; i < len(matrix); i++ {
		matrix[i] = slide(matrix[i])
	}
	return matrix
}

func slide(row []rune) []rune {
	groupStart := 0
	res := []rune{}
	for groupStart < len(row) {
		stonesCount := 0
		spaseCount := 0
	groupLoop:
		for i := groupStart; i < len(row); i++ {
			switch row[i] {
			case '#':
				for stonesCount > 0 {
					res = append(res, 'O')
					stonesCount--
				}
				for spaseCount > 0 {
					res = append(res, '.')
					spaseCount--
				}
				res = append(res, '#')
				groupStart = i + 1
				break groupLoop
			case 'O':
				stonesCount++
			case '.':
				spaseCount++
			}
			groupStart = i + 1
		}
		for stonesCount > 0 {
			res = append(res, 'O')
			stonesCount--
		}
		for spaseCount > 0 {
			res = append(res, '.')
			spaseCount--
		}
	}
	return res
}

func matrixToString(matrix [][]rune) string {
	res := strings.Builder{}
	for _, row := range matrix {
		res.WriteString(string(row))
	}
	return res.String()
}

// 90 deg
func rotateMatrix(matrix [][]rune, clockwise bool) {
	if clockwise {
		reverse(matrix)
		transpose(matrix)
	} else {
		transpose(matrix)
		reverse(matrix)

	}
}

func transpose(matrix [][]rune) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverse(matrix [][]rune) {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
}
