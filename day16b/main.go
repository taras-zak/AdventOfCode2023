package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("show me your input:")
	scanner := bufio.NewScanner(os.Stdin)

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

	var res int
	var startPoints []coord
	for i := 0; i < len(matrix[0]); i++ {
		startPoints = append(startPoints, coord{0, i, 'D'})
		startPoints = append(startPoints, coord{len(matrix) - 1, i, 'U'})
	}
	for i := 1; i < len(matrix)-1; i++ {
		startPoints = append(startPoints, coord{i, 0, 'R'})
		startPoints = append(startPoints, coord{i, len(matrix[0]) - 1, 'L'})
	}

	for _, start := range startPoints {
		visited = make(map[coord]bool)
		visitedNoDir = make(map[coord]bool)
		traverse(matrix, start)
		if len(visitedNoDir) > res {
			res = len(visitedNoDir)
		}
	}

	fmt.Println("Your answer is:", res)
}

func printMatrix(matrix [][]rune) {
	fmt.Println("")
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

type coord struct {
	row int
	col int
	dir rune
}

var visited = make(map[coord]bool)
var visitedNoDir = make(map[coord]bool)

func traverse(matrix [][]rune, curr coord) {
	if _, ok := visited[curr]; ok {
		return
	}
	visited[curr] = true
	visitedNoDir[coord{curr.row, curr.col, '.'}] = true
	for _, next := range neighbours(matrix, curr) {
		traverse(matrix, next)
	}
}

func neighbours(matrix [][]rune, start coord) []coord {
	var res []coord
	switch matrix[start.row][start.col] {
	case '.':
		// up
		if start.dir == 'U' && start.row > 0 {
			res = append(res, coord{start.row - 1, start.col, start.dir})
		}
		// down
		if start.dir == 'D' && start.row < len(matrix)-1 {
			res = append(res, coord{start.row + 1, start.col, start.dir})
		}
		// left
		if start.dir == 'L' && start.col > 0 {
			res = append(res, coord{start.row, start.col - 1, start.dir})
		}
		// right
		if start.dir == 'R' && start.col < len(matrix[0])-1 {
			res = append(res, coord{start.row, start.col + 1, start.dir})
		}
	case '/':
		// up
		if start.dir == 'U' && start.col < len(matrix[0])-1 {
			res = append(res, coord{start.row, start.col + 1, 'R'})
		}
		// down
		if start.dir == 'D' && start.col > 0 {
			res = append(res, coord{start.row, start.col - 1, 'L'})
		}
		// left
		if start.dir == 'L' && start.row < len(matrix)-1 {
			res = append(res, coord{start.row + 1, start.col, 'D'})
		}
		// right
		if start.dir == 'R' && start.row > 0 {
			res = append(res, coord{start.row - 1, start.col, 'U'})
		}
	case '\\':
		// up
		if start.dir == 'U' && start.col > 0 {
			res = append(res, coord{start.row, start.col - 1, 'L'})
		}
		// down
		if start.dir == 'D' && start.col < len(matrix[0])-1 {
			res = append(res, coord{start.row, start.col + 1, 'R'})
		}
		// left
		if start.dir == 'L' && start.row > 0 {
			res = append(res, coord{start.row - 1, start.col, 'U'})
		}
		// right
		if start.dir == 'R' && start.row < len(matrix)-1 {
			res = append(res, coord{start.row + 1, start.col, 'D'})
		}
	case '-':
		// up or down
		if start.dir == 'U' || start.dir == 'D' {
			if start.col > 0 {
				res = append(res, coord{start.row, start.col - 1, 'L'})
			}
			if start.col < len(matrix[0])-1 {
				res = append(res, coord{start.row, start.col + 1, 'R'})
			}
		}
		// left
		if start.dir == 'L' && start.col > 0 {
			res = append(res, coord{start.row, start.col - 1, start.dir})
		}
		// right
		if start.dir == 'R' && start.col < len(matrix[0])-1 {
			res = append(res, coord{start.row, start.col + 1, start.dir})
		}
	case '|':
		// up
		if start.dir == 'U' && start.row > 0 {
			res = append(res, coord{start.row - 1, start.col, start.dir})
		}
		// down
		if start.dir == 'D' && start.row < len(matrix)-1 {
			res = append(res, coord{start.row + 1, start.col, start.dir})
		}
		// left or right
		if start.dir == 'L' || start.dir == 'R' {
			if start.row > 0 {
				res = append(res, coord{start.row - 1, start.col, 'U'})
			}
			if start.row < len(matrix)-1 {
				res = append(res, coord{start.row + 1, start.col, 'D'})
			}
		}
	}
	return res
}
