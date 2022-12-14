package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2022
// Day 14, part b

func main() {
	cave := make(grid.Sparse[byte])
	ran := algo.Range[int]{-1, 1}
	maxy := 0
	for _, line := range exp.MustReadLines("inputs/14.txt") {
		pts := strings.Split(line, " -> ")
		var p image.Point
		exp.Must(fmt.Sscanf(pts[0], "%d,%d", &p.X, &p.Y))
		if p.Y > maxy {
			maxy = p.Y
		}
		for _, pt := range pts[1:] {
			var q image.Point
			exp.Must(fmt.Sscanf(pt, "%d,%d", &q.X, &q.Y))
			if q.Y > maxy {
				maxy = q.Y
			}
			d := q.Sub(p)
			d.X = ran.Clamp(d.X)
			d.Y = ran.Clamp(d.Y)
			for p != q {
				cave[p] = '#'
				p = p.Add(d)
			}
			cave[q] = '#'
		}
	}

	floor := maxy + 2

	// Sand pouring
	src := image.Pt(500, 0)
	count := 0
	for cave[src] == 0 {
		s := src
		for s.Y < floor-1 {
			fall := false
			for _, d := range []image.Point{{0, 1}, {-1, 1}, {1, 1}} {
				if t := s.Add(d); cave[t] == 0 {
					fall = true
					s = t
					break
				}
			}
			if !fall {
				break
			}
		}
		// Stopped
		count++
		cave[s] = 'O'
	}

	// c2, offset := cave.ToDense()
	// c2.Map(func(b byte) byte {
	// 	if b == 0 {
	// 		return ' '
	// 	}
	// 	return b
	// })
	// fmt.Println(offset)
	// fmt.Println(c2)

	fmt.Println(count)
}
