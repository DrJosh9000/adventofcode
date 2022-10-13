package main

import (
	"fmt"
	"image"
	"os"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2015
// Day 3, part a

func main() {
	var p image.Point
	visits := grid.Sparse[int]{{}: 1}
	move := map[byte]image.Point{
		'>': {1, 0},
		'<': {-1, 0},
		'^': {0, -1},
		'v': {0, 1},
	}
	for _, c := range exp.Must(os.ReadFile("inputs/3.txt")) {
		p = p.Add(move[c])
		visits[p]++
	}
	fmt.Println(len(visits))
}
