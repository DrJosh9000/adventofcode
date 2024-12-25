package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2024
// Day 22, part a

const inputPath = "2024/inputs/22.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	sum := 0
	for _, line := range lines {
		n := exp.MustAtoi(line)
		for range 2000 {
			n = round(n)
		}
		sum += n
	}
	fmt.Println(sum)
}

func round(n int) int {
	n ^= n << 6
	n &= 0xFFFFFF
	n ^= n >> 5
	n &= 0xFFFFFF
	n ^= n << 11
	n &= 0xFFFFFF
	return n
}
