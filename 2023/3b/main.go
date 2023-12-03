package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2023
// Day 3, part b

func main() {
	digitsRE := regexp.MustCompile(`\d+`)

	g := grid.BytesFromStrings(exp.MustReadLines("2023/inputs/3.txt"))

	stars := make(map[image.Point][]int)
	for y, row := range g {
		for x, c := range row {
			if c == '*' {
				stars[image.Pt(x, y)] = []int{}
			}
		}
	}

	for y, row := range g {
		for _, rg := range digitsRE.FindAllIndex(row, -1) {
			n := exp.Must(strconv.Atoi(string(row[rg[0]:rg[1]])))
			for s := range stars {
				if s.Y >= y-1 && s.Y <= y+1 && s.X >= rg[0]-1 && s.X <= rg[1] {
					stars[s] = append(stars[s], n)
				}
			}
		}
	}

	sum := 0
	for _, ns := range stars {
		if len(ns) != 2 {
			continue
		}
		sum += ns[0] * ns[1]
	}

	fmt.Println(sum)
}
