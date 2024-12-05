package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
)

// Advent of Code 2016
// Day 2, part b

func main() {
	keypad := map[image.Point]rune{
		{0, -2}:  '1',
		{-1, -1}: '2', {0, -1}: '3', {1, -1}: '4',
		{-2, 0}: '5', {-1, 0}: '6', {0, 0}: '7', {1, 0}: '8', {2, 0}: '9',
		{-1, 1}: 'A', {0, 1}: 'B', {1, 1}: 'C',
		{0, 2}: 'D',
	}
	dir := map[rune]image.Point{
		'U': {0, -1},
		'D': {0, 1},
		'L': {-1, 0},
		'R': {1, 0},
	}
	p := image.Pt(-2, 0)
	for _, line := range exp.MustReadLines("inputs/2.txt") {
		for _, c := range line {
			t := p.Add(dir[c])
			if _, ok := keypad[t]; ok {
				p = t
			}
		}
		fmt.Printf("%c", keypad[p])
	}
	fmt.Println()
}
