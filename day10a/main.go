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
	var maxDistance int
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
loop:
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 'S' {
				maxDistance = bfs(matrix, coord{row, col})
				break loop
			}
		}
	}
	fmt.Println("Your answer is:", maxDistance)
}

type coord struct {
	row, col int
}

func bfs(matrix [][]rune, start coord) int {
	distance := 0
	visited := map[coord]bool{}
	queue := []coord{start}
	for len(queue) > 0 {
		level := len(queue)
		for i := level; i > 0; i-- {
			top := queue[0]
			queue = queue[1:]
			if _, ok := visited[top]; ok {
				continue
			}
			visited[top] = true
			for _, n := range neighbours(matrix, top) {
				if _, ok := visited[n]; ok {
					continue
				}
				queue = append(queue, n)
			}
		}
		distance++
	}
	return distance - 1
}

var validNeighbours = map[string]map[rune]bool{
	"SU": {'|': true, 'F': true, '7': true},
	"SD": {'|': true, 'L': true, 'J': true},
	"SL": {'-': true, 'L': true, 'F': true},
	"SR": {'-': true, 'J': true, '7': true},
	"|U": {'|': true, 'F': true, '7': true},
	"|D": {'|': true, 'L': true, 'J': true},
	"-L": {'-': true, 'L': true, 'F': true},
	"-R": {'-': true, 'J': true, '7': true},
	"LU": {'|': true, 'F': true, '7': true},
	"LR": {'-': true, 'J': true, '7': true},
	"JU": {'|': true, 'F': true, '7': true},
	"JL": {'-': true, 'L': true, 'F': true},
	"7L": {'-': true, 'L': true, 'F': true},
	"7D": {'|': true, 'L': true, 'J': true},
	"FD": {'|': true, 'L': true, 'J': true},
	"FR": {'-': true, 'J': true, '7': true},
}

func neighbours(matrix [][]rune, start coord) []coord {
	var res []coord
	switch matrix[start.row][start.col] {
	case 'S':
		// up
		if start.row > 0 && validNeighbours["SU"][matrix[start.row-1][start.col]] {
			res = append(res, coord{start.row - 1, start.col})
		}
		// down
		if start.row < len(matrix)-1 && validNeighbours["SD"][matrix[start.row+1][start.col]] {
			res = append(res, coord{start.row + 1, start.col})
		}
		// left
		if start.col > 0 && validNeighbours["SL"][matrix[start.row][start.col-1]] {
			res = append(res, coord{start.row, start.col - 1})
		}
		// right
		if start.col < len(matrix[0])-1 && validNeighbours["SR"][matrix[start.row][start.col+1]] {
			res = append(res, coord{start.row, start.col + 1})
		}
	case '|':
		// up
		if start.row > 0 && validNeighbours["|U"][matrix[start.row-1][start.col]] {
			res = append(res, coord{start.row - 1, start.col})
		}
		// down
		if start.row < len(matrix)-1 && validNeighbours["|D"][matrix[start.row+1][start.col]] {
			res = append(res, coord{start.row + 1, start.col})
		}
	case '-':
		// left
		if start.col > 0 && validNeighbours["-L"][matrix[start.row][start.col-1]] {
			res = append(res, coord{start.row, start.col - 1})
		}
		// right
		if start.col < len(matrix[0])-1 && validNeighbours["-R"][matrix[start.row][start.col+1]] {
			res = append(res, coord{start.row, start.col + 1})
		}
	case 'L':
		// up
		if start.row > 0 && validNeighbours["LU"][matrix[start.row-1][start.col]] {
			res = append(res, coord{start.row - 1, start.col})
		}
		// right
		if start.col < len(matrix[0])-1 && validNeighbours["LR"][matrix[start.row][start.col+1]] {
			res = append(res, coord{start.row, start.col + 1})
		}
	case 'J':
		// up
		if start.row > 0 && validNeighbours["JU"][matrix[start.row-1][start.col]] {
			res = append(res, coord{start.row - 1, start.col})
		}
		// left
		if start.col > 0 && validNeighbours["JL"][matrix[start.row][start.col-1]] {
			res = append(res, coord{start.row, start.col - 1})
		}
	case '7':
		// down
		if start.row < len(matrix)-1 && validNeighbours["7D"][matrix[start.row+1][start.col]] {
			res = append(res, coord{start.row + 1, start.col})
		}
		// left
		if start.col > 0 && validNeighbours["7L"][matrix[start.row][start.col-1]] {
			res = append(res, coord{start.row, start.col - 1})
		}
	case 'F':
		// down
		if start.row < len(matrix)-1 && validNeighbours["FD"][matrix[start.row+1][start.col]] {
			res = append(res, coord{start.row + 1, start.col})
		}
		// right
		if start.col < len(matrix[0])-1 && validNeighbours["FR"][matrix[start.row][start.col+1]] {
			res = append(res, coord{start.row, start.col + 1})
		}
	}
	return res
}
