package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var scores []int
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
					continue lineLoop
				}
				stack = stack[:len(stack)-1]
			case ']':
				if top != '[' {
					continue lineLoop
				}
				stack = stack[:len(stack)-1]
			case '}':
				if top != '{' {
					continue lineLoop
				}
				stack = stack[:len(stack)-1]
			case '>':
				if top != '<' {
					continue lineLoop
				}
				stack = stack[:len(stack)-1]
			}
		}

		sc := 0
		for i := range stack {
			sc *= 5
			switch stack[len(stack)-i-1] {
			case '(':
				sc++
			case '[':
				sc += 2
			case '{':
				sc += 3
			case '<':
				sc += 4
			}
		}
		scores = append(scores, sc)
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
