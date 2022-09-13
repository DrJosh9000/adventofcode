package main

import (
	"fmt"
	"image"
	"strconv"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2016
// Day 1, part b

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
				fmt.Println(abs(p.X) + abs(p.Y))
				return
			}
			seen.Insert(p)
		}
	}
}
