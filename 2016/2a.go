package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
)

// Advent of Code 2016
// Day 2, part a

func main() {
	keypad := map[image.Point]int{
		{0, 0}: 1, {1, 0}: 2, {2, 0}: 3,
		{0, 1}: 4, {1, 1}: 5, {2, 1}: 6,
		{0, 2}: 7, {1, 2}: 8, {2, 2}: 9,
	}
	dir := map[rune]image.Point{
		'U': {0, -1},
		'D': {0, 1},
		'L': {-1, 0},
		'R': {1, 0},
	}
	p := image.Pt(1, 1)
	for _, line := range exp.MustReadLines("inputs/2.txt") {
		for _, c := range line {
			t := p.Add(dir[c])
			if _, ok := keypad[t]; ok {
				p = t
			}
		}
		fmt.Print(keypad[p])
	}
	fmt.Println()
}
