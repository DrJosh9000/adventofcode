package main

import (
	"fmt"
	"strconv"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2023
// Day 21, part b

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("inputs/21.txt") {
		sum += exp.Must(strconv.Atoi(line))
	}
	fmt.Println(sum)
}