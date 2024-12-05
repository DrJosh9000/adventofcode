package main

import (
	"fmt"
	"os"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2022
// Day 6, part b

func main() {
	input := exp.Must(os.ReadFile("inputs/6.txt"))

	for i := range input[13:] {
		s := algo.SetFromSlice(input[i:][:14])
		if len(s) != 14 {
			continue
		}

		fmt.Println(i + 14)
		return
	}
}
