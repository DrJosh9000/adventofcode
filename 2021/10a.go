package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/10.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	score := 0
lineLoop:
	for _, line := range lines {
		var stack []rune
		for _, r := range line {
			var top rune
			if len(stack) > 0 {
				top = stack[len(stack)-1]
			}
			switch r {
			case '(', '[', '{', '<':
				stack = append(stack, r)
			case ')':
				if top != '(' {
					score += 3
					continue lineLoop
				}
				stack = stack[:len(stack)-1]
			case ']':
				if top != '[' {
					score += 57
					continue lineLoop
				}
				stack = stack[:len(stack)-1]
			case '}':
				if top != '{' {
					score += 1197
					continue lineLoop
				}
				stack = stack[:len(stack)-1]
			case '>':
				if top != '<' {
					score += 25137
					continue lineLoop
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	fmt.Println(score)
}
