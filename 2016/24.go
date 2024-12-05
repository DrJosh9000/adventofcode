package main

import (
	"fmt"
	"image"
	"math"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2016
// Day 24

func main() {
	maze := exp.MustReadLines("inputs/24.txt")

	// Find targets in the maze.
	targmap := make(map[byte]image.Point)
	for y, line := range maze {
		for x, c := range line {
			if c >= '0' && c <= '9' {
				targmap[byte(c-'0')] = image.Pt(x, y)
			}
		}
	}
	targets, _ := algo.SliceFromMap(targmap)
	N := len(targets)

	// Find how far each target is from each other target.
	dist := grid.Make[int](N, N)
	for i, t := range targets {
		algo.FloodFill(t, func(p image.Point, d int) ([]image.Point, error) {
			c := maze[p.Y][p.X]
			if c >= '0' && c <= '9' {
				dist[i][c-'0'] = d
			}
			var next []image.Point
			for _, n := range algo.Neigh4 {
				q := p.Add(n)
				if maze[q.Y][q.X] == '#' {
					continue
				}
				next = append(next, q)
			}
			return next, nil
		})
	}

	// Try every permutation of all targets after target 0.
	// (There are very few targets...)
	s := make([]int, N)
	for i := range s {
		s[i] = i
	}
	mind, minl := math.MaxInt, math.MaxInt
	for {
		d := 0
		for i := range s[1:] {
			d += dist[s[i]][s[i+1]]
		}
		if d < mind {
			mind = d
		}
		if l := d + dist[s[N-1]][0]; l < minl {
			minl = l
		}
		if !algo.NextPermutation(s[1:]) {
			break
		}
	}

	fmt.Println(mind)
	fmt.Println(minl)
}
