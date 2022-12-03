package main

import (
	"fmt"
	"strconv"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 25

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("inputs/25.txt") {
		sum += exp.Must(strconv.Atoi(line))
	}
	fmt.Println(sum)
}