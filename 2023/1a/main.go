package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2023
// Day 1, part a

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("2023/inputs/1.txt") {
		first, last := -1, -1
		for _, d := range line {
			if d >= '0' && d <= '9' {
				if first == -1 {
					first = int(d - '0')
				}
				last = int(d - '0')
			}
		}
		sum += first*10 + last
	}
	fmt.Println(sum)
}
