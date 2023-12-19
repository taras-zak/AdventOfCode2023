package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	matrix = rotateMatrix(matrix)
	matrix = slideMatrix(matrix)
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

// -90 deg
func rotateMatrix(matrix [][]rune) [][]rune {
	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	return matrix
}
