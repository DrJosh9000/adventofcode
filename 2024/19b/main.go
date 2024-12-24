package main

import (
	"fmt"
	"regexp"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2024
// Day 19, part b

const inputPath = "2024/inputs/19.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	types := strings.Split(lines[0], ", ")
	stripesRE := regexp.MustCompile("^(" + strings.Join(types, "|") + ")+$")
	count := 0
	for _, towel := range lines[2:] {
		if !stripesRE.MatchString(towel) {
			continue
		}
		memo := map[string]int{"": 1}
		var search func(string) int
		search = func(towel string) int {
			if c, has := memo[towel]; has {
				return c
			}
			c := 0
			for _, t := range types {
				rem, ok := strings.CutPrefix(towel, t)
				if !ok {
					continue
				}
				c += search(rem)
			}
			memo[towel] = c
			return c
		}
		count += search(towel)
	}
	fmt.Println(count)
}
