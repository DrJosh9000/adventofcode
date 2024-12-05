package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2023
// Day 1, part b

var digits = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("2023/inputs/1.txt") {
		first, last := -1, -1
		for i := range line {
			for d, v := range digits {
				if strings.HasPrefix(line[i:], d) {
					if first == -1 {
						first = v
					}
					last = v
				}
			}
		}

		sum += 10*first + last
	}
	fmt.Println(sum)
}
