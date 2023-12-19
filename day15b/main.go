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
	fmt.Println("Show me your puzzle:")
	scanner := bufio.NewScanner(os.Stdin)

	var res int
	var boxes [256][]Lens

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
		for _, step := range strings.Split(line, ",") {
			substrings := strings.Split(step, "-")
			// delete
			if len(substrings) > 1 {
				label := substrings[0]
				boxIdx := hash(label)
				for lensIdx, lens := range boxes[boxIdx] {
					if lens.Label == label {
						boxes[boxIdx] = append(boxes[boxIdx][:lensIdx], boxes[boxIdx][lensIdx+1:]...)
						break
					}
				}
				continue
			}
			// add
			substrings = strings.Split(step, "=")
			label := substrings[0]
			boxIdx := hash(label)
			length, _ := strconv.Atoi(substrings[1])
			found := -1
			for lensIdx, lens := range boxes[boxIdx] {
				if lens.Label == label {
					found = lensIdx
				}
			}
			if found != -1 {
				boxes[boxIdx][found].FocusLength = length
			} else {
				boxes[boxIdx] = append(boxes[boxIdx], Lens{
					Label:       label,
					FocusLength: length,
				})
			}
		}
	}
	for i, box := range boxes {
		for j, lens := range box {
			res += (i + 1) * (j + 1) * lens.FocusLength
		}
	}
	fmt.Println("Answer:", res)
}

type Lens struct {
	Label       string
	FocusLength int
}

func hash(in string) int {
	var res int
	for _, ch := range in {
		res = ((res + int(ch)) * 17) % 256
	}
	return res
}
