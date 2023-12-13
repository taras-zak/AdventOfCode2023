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
	var stop bool
	for {
		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		if line == "" {
			if stop {
				break
			}
			stop = true
			res += findMirrors(matrix)
			matrix = nil
			continue
		}
		stop = false
		matrix = append(matrix, []rune(line))
	}

	fmt.Println("Answer:", res)
}

func findMirrors(matrix [][]rune) int {
	var res int
	rows := len(matrix)
	cols := len(matrix[0])

	// horizontal mirrors
hSplitLoop:
	for split := 0; split < rows-1; split++ {
		var diff int
		left := split
		right := split + 1
		for left >= 0 && right < rows {
			for col := 0; col < cols; col++ {
				if matrix[left][col] != matrix[right][col] {
					diff++
					if diff > 1 {
						continue hSplitLoop
					}
				}
			}
			left--
			right++
		}
		if diff == 1 {
			res += (split + 1) * 100
		}
	}

	// vertical mirrors
vSplitLoop:
	for split := 0; split < cols-1; split++ {
		var diff int
		left := split
		right := split + 1
		for left >= 0 && right < cols {
			for row := 0; row < rows; row++ {
				if matrix[row][left] != matrix[row][right] {
					diff++
					if diff > 1 {
						continue vSplitLoop
					}
				}
			}
			left--
			right++
		}
		if diff == 1 {
			res += split + 1
		}
	}
	return res
}
