package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 6

func main() {
	var comm []map[rune]int
	for _, line := range exp.MustReadLines("inputs/6.txt") {
		if len(comm) == 0 {
			comm = make([]map[rune]int, len(line))
			for i := range comm {
				comm[i] = make(map[rune]int)
			}
		}
		for i, c := range line {
			comm[i][c]++
		}
	}

	for _, h := range comm {
		c, _ := algo.MapMax(h)
		fmt.Printf("%c", c)
	}
	fmt.Println()

	for _, h := range comm {
		c, _ := algo.MapMin(h)
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
