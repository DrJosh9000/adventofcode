package main

import (
	"fmt"
	"strconv"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2022
// Day 10, part a

func main() {
	x := 1
	c := 1
	sum := 0
	want := algo.MakeSet(20, 60, 100, 140, 180, 220)
	for _, line := range exp.MustReadLines("inputs/10.txt") {
		fs := strings.Fields(line)
		if want.Contains(c) {
			sum += x * c
		}
		switch fs[0] {
		case "addx":
			if want.Contains(c + 1) {
				sum += x * (c + 1)
			}
			c += 2
			x += exp.Must(strconv.Atoi(fs[1]))
		case "noop":
			c++
		}
	}
	fmt.Println(sum)
}
