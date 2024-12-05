package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 7, part a

func main() {
	abbas := make(algo.Set[string])
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			abbas.Insert(string([]rune{a, b, b, a}))
		}
	}

	count := 0
lineLoop:
	for _, line := range exp.MustReadLines("inputs/7.txt") {
		sb, ab := false, false
		for i, c := range line {
			switch c {
			case '[':
				sb = true
			case ']':
				sb = false
			default:
				if i >= 3 && abbas.Contains(line[i-3:i+1]) {
					if sb {
						continue lineLoop
					}
					ab = true
				}
			}
		}
		if ab {
			count++
		}
	}
	fmt.Println(count)
}
