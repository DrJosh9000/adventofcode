package main

import (
	"fmt"
	"os"

	"drjosh.dev/exp"
)

// Advent of Code 2022
// Day 6, part a

func main() {
	input := exp.Must(os.ReadFile("inputs/6.txt"))

	for i := range input[3:] {
		b := input[i:][:4]
		if b[0] == b[1] || b[0] == b[2] || b[0] == b[3] || b[1] == b[2] || b[1] == b[3] || b[2] == b[3] {
			continue
		}
		fmt.Println(i + 4)
		return
	}
}
