package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 1, part a

func main() {
	max := math.MinInt
	sum := 0
	for _, line := range exp.MustReadLines("inputs/1.txt") {
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}
		sum += exp.Must(strconv.Atoi(line))
	}
	if sum > max {
		max = sum
	}
	fmt.Println(max)
}
