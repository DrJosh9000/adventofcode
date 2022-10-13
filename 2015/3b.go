package main

import (
	"fmt"
	"image"
	"os"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2015
// Day 3, part b

func main() {
	p := []image.Point{{}, {}}
	visits := grid.Sparse[int]{{}: 1}
	move := map[byte]image.Point{
		'>': {1, 0},
		'<': {-1, 0},
		'^': {0, -1},
		'v': {0, 1},
	}
	for i, c := range exp.Must(os.ReadFile("inputs/3.txt")) {
		t := i % 2
		p[t] = p[t].Add(move[c])
		visits[p[t]]++
	}
	fmt.Println(len(visits))
}
