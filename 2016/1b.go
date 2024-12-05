package main

import (
	"fmt"
	"image"
	"strconv"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 1, part b

func main() {
	seen := make(algo.Set[image.Point])
	p, d := image.Pt(0, 0), image.Pt(0, -1)
	seen.Insert(p)
	for _, step := range exp.MustReadDelimited("inputs/1.txt", ",") {
		if step[0] == 'L' {
			d = image.Pt(d.Y, -d.X)
		} else {
			d = image.Pt(-d.Y, d.X)
		}
		for i := 0; i < exp.Must(strconv.Atoi(step[1:])); i++ {
			p = p.Add(d)
			if seen.Contains(p) {
				fmt.Println(algo.Abs(p.X) + algo.Abs(p.Y))
				return
			}
			seen.Insert(p)
		}
	}
}
