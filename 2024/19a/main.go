package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2024
// Day 19, part a

const inputPath = "2024/inputs/19.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	sum := 0
	for _, line := range lines {
		sum += exp.MustAtoi(line)
	}
	fmt.Println(sum)
}