package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2023
// Day 6, part b

func main() {
	lines := exp.MustReadLines("2023/inputs/6.txt")

	t := exp.Must(strconv.Atoi(strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time:"), " ", "")))
	d := exp.Must(strconv.Atoi(strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance:"), " ", "")))

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

	fmt.Println(maxs - mins + 1)

}
