package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type nodeNeighbors struct {
	left  string
	right string
}

func main() {
	fmt.Println("show me your input:")
	scanner := bufio.NewScanner(os.Stdin)

	var res int
	var matrix [][]int
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

		nodes := strings.Split(line, " ")

		row := []int{0}
		for _, node := range nodes {
			d, _ := strconv.Atoi(node)
			row = append(row, d)
		}
		matrix = append(matrix, row)

	levelLoop:
		for level := 0; ; level++ {
			row = []int{0}
			for i := 1; i < len(matrix[level])-1; i++ {
				row = append(row, matrix[level][i+1]-matrix[level][i])
			}

			for j := 1; j < len(row)-1; j++ {
				if row[j] != row[j+1] {
					matrix = append(matrix, row)
					continue levelLoop
				}
			}
			row[0] = row[1]
			matrix = append(matrix, row)
			break levelLoop
		}
		for i := len(matrix) - 2; i >= 0; i-- {
			matrix[i][0] = matrix[i][1] - matrix[i+1][0]
		}
		res += matrix[0][0]
		fmt.Println("matrix:", matrix)
		matrix = nil
	}
	fmt.Println("Your answer is:", res)
}
