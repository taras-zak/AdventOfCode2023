package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	row, col, direction, steps, heatLoss int
}

// PriorityQueue is a min-heap implementation for Dijkstra's algorithm.
type PriorityQueue []Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].heatLoss < pq[j].heatLoss }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

type visitedNode struct {
	row, col  int
	direction int
	steps     int
}

func findMinHeatLoss(heatMap [][]int) int {
	rows := len(heatMap)
	cols := len(heatMap[0])

	// Define directions: right, down, left, up.
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	// Create a 2D array to keep track of visited blocks.
	visited := make(map[visitedNode]int)

	// Priority queue for Dijkstra's algorithm.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// From top left go right and down
	heap.Push(&pq, Node{0, 1, 0, 1, 0})
	heap.Push(&pq, Node{1, 0, 1, 1, 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(Node)
		if current.row >= rows || current.col >= cols || current.col < 0 || current.row < 0 {
			continue
		}
		heatLoss := current.heatLoss + heatMap[current.row][current.col]

		// Check if we reached the destination.
		if current.row == rows-1 && current.col == cols-1 {
			return heatLoss
		}

		if val, ok := visited[visitedNode{current.row, current.col, current.direction, current.steps}]; ok {
			if val <= heatLoss {
				continue
			}
		}

		visited[visitedNode{current.row, current.col, current.direction, current.steps}] = heatLoss

		// turn right
		dir := mod(current.direction+1, 4)
		heap.Push(&pq, Node{current.row + dx[dir], current.col + dy[dir], dir, 1, heatLoss})
		// turn left
		dir = mod(current.direction-1, 4)
		heap.Push(&pq, Node{current.row + dx[dir], current.col + dy[dir], dir, 1, heatLoss})

		// move straight
		if current.steps < 3 {
			heap.Push(&pq, Node{current.row + dx[current.direction], current.col + dy[current.direction], current.direction, current.steps + 1, heatLoss})
		}
	}

	// If no valid path is found.
	return -1
}

// divmod that also handles negative i
func mod(i, n int) int {
	return ((i % n) + n) % n
}

func main() {
	fmt.Println("show me your input:")
	scanner := bufio.NewScanner(os.Stdin)

	var heatMap [][]int
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

		nodes := strings.Split(line, "")
		var row []int
		for _, node := range nodes {
			d, _ := strconv.Atoi(node)
			row = append(row, d)
		}
		heatMap = append(heatMap, row)
	}

	result := findMinHeatLoss(heatMap)
	fmt.Println("Answer:", result)
}
