package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 7, part b

func main() {
	abas := make(map[string]string)
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			abas[string([]rune{a, b, a})] = string([]rune{b, a, b})
		}
	}

	count := 0
	for _, line := range exp.MustReadLines("inputs/7.txt") {
		abs, bas := make(algo.Set[string]), make(algo.Set[string])
		sb := false
		for i, c := range line {
			switch c {
			case '[':
				sb = true
			case ']':
				sb = false
			default:
				if i < 2 {
					continue
				}
				aba := line[i-2 : i+1]
				if bab := abas[aba]; bab != "" {
					if sb {
						bas.Insert(bab)
					} else {
						abs.Insert(aba)
					}
				}
			}
		}
		if !abs.Disjoint(bas) {
			count++
		}
	}
	fmt.Println(count)
}
