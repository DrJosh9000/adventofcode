package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2024
// Day 25

const inputPath = "2024/inputs/25.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	var locks, keys [][5]int
lineLoop:
	for i := 0; i < len(lines); i += 8 {
		var hs [5]int
		for j := range 5 {
			if i+j+1 >= len(lines) {
				break lineLoop
			}
			row := lines[i+j+1]
			for k, c := range row {
				if c == '#' {
					hs[k]++
				}
			}
		}
		if lines[i] == "#####" {
			locks = append(locks, hs)
		} else {
			keys = append(keys, hs)
		}
	}
	count := 0
	for _, l := range locks {
		for _, k := range keys {
			overlap := false
			for i := range 5 {
				if l[i]+k[i] > 5 {
					overlap = true
					break
				}
			}
			if !overlap {
				count++
			}
		}
	}

	fmt.Println(count)
}
