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

	var time int
	var distance int
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
			time, _ = strconv.Atoi(strings.Replace(v, " ", "", -1))
		case "Distance":
			distance, _ = strconv.Atoi(strings.Replace(v, " ", "", -1))
		}
	}
	result := 1
	var count int
	for holdDuration := 1; holdDuration < time; holdDuration++ {
		dist := holdDuration * (time - holdDuration)
		if dist > distance {
			count++
		}
	}
	result *= count

	fmt.Println("Your answer is: ", result)
}
