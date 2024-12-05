package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 5, part b

func main() {
	nice := 0
	for _, line := range exp.MustReadLines("inputs/5.txt") {
		r := false
		for i := range line[2:] {
			if line[i] == line[i+2] {
				r = true
				break
			}
		}
		if !r {
			continue
		}

		p := false
	outerLoop:
		for i := range line[3:] {
			for j := range line[i+3:] {
				if line[i] == line[i+j+2] && line[i+1] == line[i+j+3] {
					p = true
					break outerLoop
				}
			}
		}
		if !p {
			continue
		}
		nice++
	}
	fmt.Println(nice)
}
