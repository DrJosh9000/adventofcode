package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 6, part a

func main() {
	lines := exp.MustReadLines("2023/inputs/6.txt")
	times := algo.MustMap(strings.Fields(lines[0])[1:], strconv.Atoi)
	dists := algo.MustMap(strings.Fields(lines[1])[1:], strconv.Atoi)

	prod := 1
	for i, t := range times {
		d := dists[i]

		mins := 0
		for ; mins < t; mins++ {
			if mins*(t-mins) > d {
				break
			}
		}
		maxs := t
		for ; maxs >= 0; maxs-- {
			if maxs*(t-maxs) > d {
				break
			}
		}

		prod *= (maxs - mins + 1)
	}

	fmt.Println(prod)
}
