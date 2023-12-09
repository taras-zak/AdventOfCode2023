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

	var startNodes []string
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
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
		neighbours := strings.Split(strings.Trim(strings.Split(line, " = ")[1], "()"), ", ")
		nodesMap[node] = nodeNeighbors{left: neighbours[0], right: neighbours[1]}
	}

	var stepsArr []int
	for _, curr := range startNodes {
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
				if strings.HasSuffix(curr, "Z") {
					break directionsCycle
				}
			}
		}
		stepsArr = append(stepsArr, steps)
	}

	fmt.Println("Your answer is:", lcm(stepsArr[0], stepsArr[1], stepsArr[2:]...))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
