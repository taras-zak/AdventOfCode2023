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

	var times []int
	var distances []int
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		k := strings.Split(line, ":")[0]
		v := strings.Split(line, ":")[1]
		switch k {
		case "Time":
			for _, part := range strings.Split(v, " ") {
				if part != "" {
					time, _ := strconv.Atoi(part)
					times = append(times, time)
				}
			}
		case "Distance":
			for _, part := range strings.Split(v, " ") {
				if part != "" {
					distance, _ := strconv.Atoi(part)
					distances = append(distances, distance)
				}
			}
		}
	}
	result := 1
	for i := 0; i < len(times); i++ {
		var count int
		for holdDuration := 1; holdDuration < times[i]; holdDuration++ {
			dist := holdDuration * (times[i] - holdDuration)
			if dist > distances[i] {
				count++
			}
		}
		result *= count
	}
	fmt.Println("Your answer is: ", result)
}
