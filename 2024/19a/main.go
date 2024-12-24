package main

import (
	"fmt"
	"regexp"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2024
// Day 19, part a

const inputPath = "2024/inputs/19.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	types := strings.Split(lines[0], ", ")
	stripesRE := regexp.MustCompile("^(" + strings.Join(types, "|") + ")+$")
	count := 0
	for _, towel := range lines[2:] {
		if stripesRE.MatchString(towel) {
			count++
		}
	}
	fmt.Println(count)
}
