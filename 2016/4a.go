package main

import (
	"fmt"
	"strconv"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"golang.org/x/exp/maps"
)

// Advent of Code 2016
// Day 4, part a

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("inputs/4.txt") {
		lc := make(map[rune]int)
		parts := strings.FieldsFunc(line, func(c rune) bool {
			return c == '-' || c == '[' || c == ']'
		})
		for _, part := range parts[:len(parts)-2] {
			for _, c := range part {
				lc[c]++
			}
		}
		cs := maps.Keys(lc)
		algo.SortAsc(cs)
		algo.SortByMapDesc(cs, lc)
		if string(cs[:5]) != parts[len(parts)-1] {
			continue
		}
		sum += exp.Must(strconv.Atoi(parts[len(parts)-2]))
	}
	fmt.Println(sum)
}
