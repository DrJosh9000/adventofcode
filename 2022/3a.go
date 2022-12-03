package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 3, part a

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("inputs/3.txt") {
		l, r := make(algo.Set[rune]), make(algo.Set[rune])
		for i, c := range line {
			if i < len(line)/2 {
				l.Insert(c)
			} else {
				r.Insert(c)
			}
		}
		for c := range l.Intersection(r) {
			if c > 'a' {
				sum += int(c) - 'a' + 1
			} else {
				sum += int(c) - 'A' + 27
			}
		}
	}
	fmt.Println(sum)
}
