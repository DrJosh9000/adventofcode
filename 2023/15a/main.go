package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2023
// Day 15, part a

const inputPath = "2023/inputs/15.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	sum := 0
	for _, line := range lines {
		for _, token := range strings.Split(line, ",") {
			h := 0
			for _, c := range token {
				h += int(c)
				h *= 17
				h %= 256
			}
			sum += h
		}
	}
	fmt.Println(sum)
}
