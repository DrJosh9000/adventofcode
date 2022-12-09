package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 9, part a

func norm(p image.Point) int {
	return algo.Max(algo.Abs(p.X), algo.Abs(p.Y))
}

func main() {
	visited := make(algo.Set[image.Point])
	var h, t image.Point
	for _, line := range exp.MustReadLines("inputs/9.txt") {
		var d rune
		var s int
		exp.Must(fmt.Sscanf(line, "%c %d", &d, &s))
		for i := 0; i < s; i++ {
			switch d {
			case 'R':
				h = h.Add(image.Pt(1, 0))
			case 'L':
				h = h.Add(image.Pt(-1, 0))
			case 'U':
				h = h.Add(image.Pt(0, -1))
			case 'D':
				h = h.Add(image.Pt(0, 1))
			}
			delta := h.Sub(t)
			if norm(delta) > 1 {
				step := delta
				if delta.X < -1 {
					step.X = -1
				}
				if delta.X > 1 {
					step.X = 1
				}
				if delta.Y < -1 {
					step.Y = -1
				}
				if delta.Y > 1 {
					step.Y = 1
				}
				t = t.Add(step)
			}

			visited.Insert(t)
		}
	}
	fmt.Println(len(visited))
}
