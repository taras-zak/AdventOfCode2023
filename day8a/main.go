package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type nodeNeighbors struct {
	left  string
	right string
}

func main() {
	fmt.Println("show me your input:")
	scanner := bufio.NewScanner(os.Stdin)

	root := "AAA"
	nodesMap := make(map[string]nodeNeighbors)

	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()
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

		node := strings.Split(line, " = ")[0]
		neighbours := strings.Split(strings.Trim(strings.Split(line, " = ")[1], "()"), ", ")
		nodesMap[node] = nodeNeighbors{left: neighbours[0], right: neighbours[1]}
	}

	curr := root
	steps := 0
directionsCycle:
	for {
		for _, dir := range directions {
			steps++
			switch dir {
			case 'R':
				curr = nodesMap[curr].right
			case 'L':
				curr = nodesMap[curr].left
			}
			if curr == "ZZZ" {
				break directionsCycle
			}
		}
	}
	fmt.Println("Your answer is:", steps)
}
