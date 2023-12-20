package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Show me your input:")
	scanner := bufio.NewScanner(os.Stdin)

	var plan []*DigMove
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
		action, _ := parseDigMove(line)
		plan = append(plan, action)
	}

	fmt.Println("Your answer is:", digLagoon(plan))
}

type Point struct {
	X, Y int
}

func digLagoon(digPlan []*DigMove) int {
	var points []Point
	var perimeter int

	directions := map[string][2]int{"U": {-1, 0}, "D": {1, 0}, "L": {0, -1}, "R": {0, 1}}

	// Current position
	currentPosition := Point{0, 0}

	for _, move := range digPlan {
		points = append(points, currentPosition)
		currentPosition.X += directions[move.Direction][0] * move.Distance
		currentPosition.Y += directions[move.Direction][1] * move.Distance
		perimeter += move.Distance
	}
	return calculatePolygonArea(points) + perimeter/2.0 + 1
}

func calculatePolygonArea(vertices []Point) int {
	n := len(vertices)
	area := 0

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += (vertices[i].X + vertices[j].X) * (vertices[i].Y - vertices[j].Y)
	}

	area = area / 2.0
	return area
}

type DigMove struct {
	Direction string
	Distance  int
	Color     string
}

func parseDigMove(move string) (*DigMove, error) {
	// Define a regular expression to match the pattern
	re := regexp.MustCompile(`^([UDLR]) (\d+) \(\#([0-9a-fA-F]{6})\)$`)

	// Use the regular expression to extract components
	match := re.FindStringSubmatch(move)
	if match == nil {
		return nil, fmt.Errorf("invalid move format: %s", move)
	}

	// Parse the components
	direction := match[1]
	distance := match[2]
	color := match[3]

	// Convert distance to an integer
	var dist int
	_, err := fmt.Sscanf(distance, "%d", &dist)
	if err != nil {
		return nil, fmt.Errorf("failed to parse distance: %s", err)
	}

	// Create and return a DigMove struct
	return &DigMove{
		Direction: direction,
		Distance:  dist,
		Color:     color,
	}, nil
}
