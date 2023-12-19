package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Show me your puzzle:")
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
		for _, group := range strings.Split(line, ",") {
			var curr int
			for _, ch := range group {
				curr = ((curr + int(ch)) * 17) % 256
			}
			res += curr
		}
	}
	fmt.Println("Answer:", res)
}
