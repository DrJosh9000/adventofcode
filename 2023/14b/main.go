package main

import (
	"fmt"
	"strconv"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2023
// Day 14, part b

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("2023/inputs/14.txt") {
		sum += exp.Must(strconv.Atoi(line))
	}
	fmt.Println(sum)
}