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
		groups := []int{}
		for _, groupString := range strings.Split(strings.Split(line, " ")[1], ",") {
			group, _ := strconv.Atoi(groupString)
			groups = append(groups, group)
		}
		pattern, groups = unfold(pattern, groups)
		cache = map[entry]int{}
		res += bf(pattern, groups, -1)
	}
	fmt.Println("Answer:", res)
}

func unfold(pattern []rune, groups []int) ([]rune, []int) {
	patternUnfolded := pattern
	groupsUnfolded := groups
	for i := 0; i < 4; i++ {
		patternUnfolded = append(patternUnfolded, '?')
		patternUnfolded = append(patternUnfolded, pattern...)
		groupsUnfolded = append(groupsUnfolded, groups...)
	}
	return patternUnfolded, groupsUnfolded
}

var cache map[entry]int

type entry struct {
	patternLen  int
	groupsLen   int
	leftInGroup int
}

func bf(pattern []rune, groups []int, leftInGroup int) int {
	if 0 == len(pattern) {
		if len(groups) == 0 || (len(groups) == 1 && leftInGroup == 0) {
			return 1
		}
		return 0
	}
	if cached, ok := cache[entry{len(pattern), len(groups), leftInGroup}]; ok {
		return cached
	}
	var res int
	curr := pattern[0]
	switch curr {
	case '?':
		if leftInGroup == 0 {
			res = bf(pattern[1:], groups[1:], -1)
		} else if leftInGroup == -1 {
			if len(groups) > 0 {
				res += bf(pattern[1:], groups, groups[0]-1)
			}
			res += bf(pattern[1:], groups, -1)
		} else {
			res += bf(pattern[1:], groups, leftInGroup-1)
		}
	case '.':
		if leftInGroup == 0 {
			res = bf(pattern[1:], groups[1:], -1)
		} else if leftInGroup == -1 {
			res = bf(pattern[1:], groups, leftInGroup)
		}
	case '#':
		if leftInGroup == -1 && len(groups) > 0 {
			res = bf(pattern[1:], groups, groups[0]-1)
		} else if leftInGroup > 0 {
			res = bf(pattern[1:], groups, leftInGroup-1)
		}
	}
	cache[entry{len(pattern), len(groups), leftInGroup}] = res
	return res
}
