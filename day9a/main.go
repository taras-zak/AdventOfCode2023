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

		row := []int{}
		for _, node := range nodes {
			d, _ := strconv.Atoi(node)
			row = append(row, d)
		}
		matrix = append(matrix, row)

	levelLoop:
		for level := 0; ; level++ {
			row = []int{}
			for i := 0; i < len(matrix[level])-1; i++ {
				row = append(row, matrix[level][i+1]-matrix[level][i])
			}

			for j := 0; j < len(row)-1; j++ {
				if row[j] != row[j+1] {
					matrix = append(matrix, row)
					continue levelLoop
				}
			}
			matrix = append(matrix, append(row, row[0]))
			break levelLoop
		}
		for i := len(matrix) - 2; i >= 0; i-- {
			matrix[i] = append(matrix[i], matrix[i][len(matrix[i])-1]+matrix[i+1][len(matrix[i+1])-1])
		}
		res += matrix[0][len(matrix[0])-1]
		fmt.Println("matrix:", matrix)
		matrix = nil
	}
	fmt.Println("Your answer is:", res)
}
