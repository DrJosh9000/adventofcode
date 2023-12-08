package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2023
// Day 15, part a

const inputPath = "2023/inputs/15.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	sum := 0
	for _, line := range lines {
		sum += exp.MustAtoi(line)
	}
	fmt.Println(sum)
}