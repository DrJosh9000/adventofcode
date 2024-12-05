package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2022
// Day 9, part a

func main() {
	visited := make(algo.Set[image.Point])
	ran := algo.Range[int]{-1, 1}
	var h, t image.Point
	for _, line := range exp.MustReadLines("inputs/9.txt") {
		var d rune
		var s int
		exp.Must(fmt.Sscanf(line, "%c %d", &d, &s))
		for i := 0; i < s; i++ {
			h = h.Add(algo.ULDR[d])
			delta := h.Sub(t)
			if algo.Linfty(delta) > 1 {
				step := delta
				step.X = ran.Clamp(step.X)
				step.Y = ran.Clamp(step.Y)
				t = t.Add(step)
			}

			visited.Insert(t)
		}
	}
	fmt.Println(len(visited))
}
