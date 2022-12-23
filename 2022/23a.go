package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2022
// Day 23, part a

func main() {
	gd := exp.MustReadByteGrid("inputs/23.txt")
	elves := grid.MapToSparse(gd, func(b byte) (bool, bool) {
		return true, b == '#'
	})

	// fmt.Println(elves.ToDense())

	step := []*struct {
		name string
		e    []image.Point
		s    image.Point
	}{
		{"alone", algo.Neigh8, image.Point{}},
		{"north", []image.Point{{0, -1}, {1, -1}, {-1, -1}}, image.Point{0, -1}}, // n, ne, nw => n
		{"south", []image.Point{{0, 1}, {1, 1}, {-1, 1}}, image.Point{0, 1}},     // s, se, sw => s
		{"west", []image.Point{{-1, 0}, {-1, -1}, {-1, 1}}, image.Point{-1, 0}},  // w, nw, sw => w
		{"east", []image.Point{{1, 0}, {1, -1}, {1, 1}}, image.Point{1, 0}},      // e, ne, se => e
	}

	for round := 0; round < 10; round++ {
		// log.Printf("Round %d", round)
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

		e2 := make(grid.Sparse[bool])
		for e, p := range prop {
			if pc[p] == 1 {
				e2[p] = true
			} else {
				e2[e] = true
			}
		}
		elves = e2

		// fmt.Println(elves.ToDense())

		step[1], step[2], step[3], step[4] = step[2], step[3], step[4], step[1]
		// for _, s := range step {
		// 	fmt.Println(s.name)
		// }
	}

	bounds := elves.Bounds()
	fmt.Println(bounds.Dx()*bounds.Dy() - len(elves))
}
