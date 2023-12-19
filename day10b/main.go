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
	var theLoopMap map[coord]int
loop:
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 'S' {
				theLoopMap = traverse(matrix, coord{row, col}, 0)
				break loop
			}
		}
	}
	findStartSymbol(matrix, loopPath)
	fmt.Println(countInsideLoop(matrix, theLoopMap))
}

func countInsideLoop(matrix [][]rune, theLoopMap map[coord]int) int {
	var res int
	for row := range matrix {
		for col := range matrix[row] {
			curr := coord{row, col}
			if isInside(matrix, theLoopMap, curr) {
				res++
			}

		}
	}
	return res
}

func isInside(matrix [][]rune, theLoopMap map[coord]int, curr coord) bool {
	if _, ok := theLoopMap[curr]; ok {
		return false
	}
	count := 0
	var lastChar rune
	for {
		curr = coord{curr.row, curr.col + 1}
		if curr.col >= len(matrix[0]) {
			return count%2 != 0
		}
		if _, ok := theLoopMap[curr]; ok {
			switch matrix[curr.row][curr.col] {
			case 'F', '7', 'J', 'L':
				if lastChar == 0 {
					count++
					lastChar = matrix[curr.row][curr.col]
				} else {
					if lastChar == 'L' && matrix[curr.row][curr.col] != '7' {
						count++
					}
					if lastChar == 'F' && matrix[curr.row][curr.col] != 'J' {
						count++
					}
					lastChar = 0
				}
			case '|':
				count++
			case '-':
			default:
				panic("not reachable")
			}
		}
	}
}
func findStartSymbol(matrix [][]rune, path []coord) {
	start := path[0]
	last := path[len(path)-1]
	second := path[1]
	for _, possibleChar := range "|-FLJ7" {
		theLoopMap := make(map[coord]int)
		theLoopMap[path[len(path)-2]] = -1
		matrix[start.row][start.col] = possibleChar
		printMatrix(matrix)
		if getNext(matrix, last, theLoopMap) != start {
			continue
		}
		if getNext(matrix, start, theLoopMap) != second {
			continue
		}
		return
	}
	panic("not reachable")
}

func printMatrix(matrix [][]rune) {
	fmt.Println("")
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

type coord struct {
	row, col int
}

var loopPath []coord

func traverse(matrix [][]rune, start coord, counter int) map[coord]int {
	theLoopMap := make(map[coord]int)
	curr := start
	for {
		theLoopMap[curr] = counter
		loopPath = append(loopPath, curr)
		//prev := curr
		curr = getNext(matrix, curr, theLoopMap)
		//matrix[prev.row][prev.col] = rune('0' + counter)
		//printMatrix(matrix)
		if curr.row == -1 && curr.col == -1 {
			break
		}
		counter++
	}
	return theLoopMap
}

var validNeighbours = map[rune]map[rune]bool{
	'U': {'|': true, 'F': true, '7': true},
	'D': {'|': true, 'L': true, 'J': true},
	'L': {'-': true, 'L': true, 'F': true},
	'R': {'-': true, 'J': true, '7': true},
}
var vn = map[rune]string{
	'|': "UD",
	'-': "RL",
	'F': "DR",
	'7': "LD",
	'L': "RU",
	'J': "LU",
}

func getDirectionForStart(matrix [][]rune, start coord) coord {
	if start.row > 0 && validNeighbours['U'][matrix[start.row-1][start.col]] {
		return coord{start.row - 1, start.col}
	}
	if start.row < len(matrix)-1 && validNeighbours['D'][matrix[start.row+1][start.col]] {
		return coord{start.row + 1, start.col}
	}
	if start.col > 0 && validNeighbours['L'][matrix[start.row][start.col-1]] {
		return coord{start.row, start.col - 1}
	}
	if start.col < len(matrix[0])-1 && validNeighbours['R'][matrix[start.row][start.col+1]] {
		return coord{start.row, start.col + 1}
	}
	panic("not reachable")
}

func getNext(matrix [][]rune, curr coord, theLoopMap map[coord]int) coord {
	currChar := matrix[curr.row][curr.col]
	if currChar == 'S' {
		return getDirectionForStart(matrix, curr)
	}
	for _, dir := range vn[currChar] {
		switch dir {
		case 'U':
			next := coord{curr.row - 1, curr.col}
			if _, ok := theLoopMap[next]; curr.row > 0 && !ok {
				return next
			}
		case 'D':
			next := coord{curr.row + 1, curr.col}
			if _, ok := theLoopMap[next]; curr.row < len(matrix)-1 && !ok {
				return next
			}
		case 'L':
			next := coord{curr.row, curr.col - 1}
			if _, ok := theLoopMap[next]; curr.col > 0 && !ok {
				return next
			}
		case 'R':
			next := coord{curr.row, curr.col + 1}
			if _, ok := theLoopMap[next]; curr.col < len(matrix[0])-1 && !ok {
				return next
			}
		}
	}
	return coord{-1, -1}
}
