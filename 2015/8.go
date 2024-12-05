package main

import (
	"fmt"
	"strconv"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 8

func main() {
	suma, sumb := 0, 0
	for _, line := range exp.MustReadLines("inputs/8.txt") {
		suma += len(line) - len(exp.Must(strconv.Unquote(line)))
		sumb += len(strconv.Quote(line)) - len(line)
	}
	fmt.Println(suma, sumb)
}
