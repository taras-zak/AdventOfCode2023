package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Show me your universe:")
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

	rowsOccupied := map[int]bool{}
	colsOccupied := map[int]bool{}
	var galaxiesLoc []coord
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '#' {
				rowsOccupied[row] = true
				colsOccupied[col] = true
				galaxiesLoc = append(galaxiesLoc, coord{row, col})
			}
		}
	}

	res1 := 0
	res2 := 0
	for gal1 := 0; gal1 < len(galaxiesLoc); gal1++ {
		for gal2 := gal1 + 1; gal2 < len(galaxiesLoc); gal2++ {
			res1 += galaxiesLoc[gal1].distance(galaxiesLoc[gal2], rowsOccupied, colsOccupied, 1)
			res2 += galaxiesLoc[gal1].distance(galaxiesLoc[gal2], rowsOccupied, colsOccupied, 999999)
		}
	}

	fmt.Println("Answer for part 1:", res1)
	fmt.Println("Answer for part 2:", res2)
}

type coord struct {
	row, col int
}

func (c coord) distance(another coord, rowsOccupied, colsOccupied map[int]bool, additive int) int {
	x1 := c.col
	x2 := another.col
	if x1 < x2 {
		x1, x2 = x2, x1
	}
	dx := x1 - x2
	for i := x2 + 1; i < x1; i++ {
		if !colsOccupied[i] {
			dx += additive
		}
	}

	y1 := c.row
	y2 := another.row
	if y1 < y2 {
		y1, y2 = y2, y1
	}
	dy := y1 - y2
	for i := y2 + 1; i < y1; i++ {
		if !rowsOccupied[i] {
			dy += additive
		}
	}
	return dx + dy
}
