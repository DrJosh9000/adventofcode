package main

import (
	"fmt"
	"strconv"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2022
// Day 1, part b

func main() {
	var elves []int
	sum := 0
	for _, line := range exp.MustReadLines("inputs/1.txt") {
		if line == "" {
			elves = append(elves, sum)
			sum = 0
			continue
		}
		sum += exp.Must(strconv.Atoi(line))
	}
	elves = append(elves, sum)
	algo.SortDesc(elves)
	fmt.Println(algo.Sum(elves[:3]))
}
