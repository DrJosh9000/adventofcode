package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2022
// Day 23, part b

func main() {
	gd := exp.MustReadByteGrid("inputs/23.txt")
	elves := grid.MapToSparse(gd, func(b byte) (bool, bool) {
		return true, b == '#'
	})

	step := []*struct {
		e []image.Point
		s image.Point
	}{
		{algo.Neigh8, image.Point{}},
		{[]image.Point{{0, -1}, {1, -1}, {-1, -1}}, image.Point{0, -1}},
		{[]image.Point{{0, 1}, {1, 1}, {-1, 1}}, image.Point{0, 1}},
		{[]image.Point{{-1, 0}, {-1, -1}, {-1, 1}}, image.Point{-1, 0}},
		{[]image.Point{{1, 0}, {1, -1}, {1, 1}}, image.Point{1, 0}},
	}

	for round := 0; ; round++ {
		prop := make(map[image.Point]image.Point)
		pc := make(map[image.Point]int)
		for e := range elves {
			prop[e] = e
		stepLoop:
			for _, s := range step {
				for _, d := range s.e {
					if elves[e.Add(d)] {
						continue stepLoop
					}
				}
				p := e.Add(s.s)
				prop[e] = p
				pc[p]++
				break
			}
		}

		moved := false
		for e, p := range prop {
			if pc[p] == 1 && e != p {
				moved = true
				break
			}
		}
		if !moved {
			fmt.Println(round + 1)
			return
		}

		e2 := make(grid.Sparse[bool])
		for e, p := range prop {
			if pc[p] == 1 {
				e2[p] = true
			} else {
				e2[e] = true
			}
		}
		elves = e2

		step[1], step[2], step[3], step[4] = step[2], step[3], step[4], step[1]
	}
}
