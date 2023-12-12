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
	fmt.Println("Show me your universe:")
	scanner := bufio.NewScanner(os.Stdin)
	var res int
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

		pattern := []rune(strings.Split(line, " ")[0])
		groups := strings.Split(strings.Split(line, " ")[1], ",")
		res += bf(pattern, groups, 0)
	}
	fmt.Println("Answer:", res)
}

func bf(pattern []rune, groups []string, currPos int) int {
	if currPos == len(pattern) {
		if isValid(pattern, groups) {
			return 1
		}
		return 0
	}
	if pattern[currPos] == '?' {
		changedDot := make([]rune, len(pattern))
		copy(changedDot, pattern)
		changedDot[currPos] = '.'
		changedSharp := make([]rune, len(pattern))
		copy(changedSharp, pattern)
		changedSharp[currPos] = '#'
		return bf(changedDot, groups, currPos+1) + bf(changedSharp, groups, currPos+1)
	}
	return bf(pattern, groups, currPos+1)
}

func isValid(pattern []rune, groups []string) bool {
	res := []string{}
	groupLen := 0
	for i, ch := range pattern {
		switch ch {
		case '#':
			groupLen++
			if i == len(pattern)-1 {
				res = append(res, strconv.Itoa(groupLen))
			}
		case '.':
			if groupLen > 0 {
				res = append(res, strconv.Itoa(groupLen))
			}
			groupLen = 0
		}
	}
	if len(res) != len(groups) {
		return false
	}
	for idx := range groups {
		if res[idx] != groups[idx] {
			return false
		}
	}
	return true
}
