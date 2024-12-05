package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2022
// Day 9, part b

func main() {
	visited := make(algo.Set[image.Point])
	ran := algo.Range[int]{-1, 1}
	snake := make([]image.Point, 10)
	for _, line := range exp.MustReadLines("inputs/9.txt") {
		var d rune
		var s int
		exp.Must(fmt.Sscanf(line, "%c %d", &d, &s))
		for i := 0; i < s; i++ {
			snake[0] = snake[0].Add(algo.ULDR[d])
			for j := range snake[1:] {
				delta := snake[j].Sub(snake[j+1])
				if algo.Linfty(delta) <= 1 {
					continue
				}
				step := delta
				step.X = ran.Clamp(step.X)
				step.Y = ran.Clamp(step.Y)
				snake[j+1] = snake[j+1].Add(step)
			}

			visited.Insert(snake[9])
		}
	}
	fmt.Println(len(visited))
}
