package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2015
// Day 18, part a

func main() {
	g := make(algo.Set[image.Point])
	for y, row := range exp.MustReadLines("inputs/18.txt") {
		for x, c := range row {
			if c == '#' {
				g.Insert(image.Pt(x, y))
			}
		}
	}

	for i := 0; i < 100; i++ {
		g2 := make(algo.Set[image.Point])
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				p := image.Pt(x, y)
				c := 0
				for _, d := range algo.Neigh8 {
					if g.Contains(p.Add(d)) {
						c++
					}
				}
				if c == 3 || (g.Contains(p) && c == 2) {
					g2.Insert(p)
				}
			}
		}
		g = g2
	}

	fmt.Println(len(g))
}
