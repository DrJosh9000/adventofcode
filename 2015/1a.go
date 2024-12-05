package main

import (
	"fmt"
	"os"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2015
// Day 1, part a

func main() {
	h := algo.Freq(exp.Must(os.ReadFile("inputs/1.txt")))
	fmt.Println(h['('] - h[')'])
}
