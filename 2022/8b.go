package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 8, part b

func main() {
	g := exp.MustReadByteGrid("inputs/8.txt")

	bounds := g.Bounds()

	maxscore := 0
	for r := range g {
		for c, h := range g[r] {
			score := 1
			for _, d := range algo.Neigh4 {
				dist := 0
				p := image.Pt(c, r).Add(d)
				for p.In(bounds) {
					dist++
					if g[p.Y][p.X] >= h {
						break
					}
					p = p.Add(d)
				}
				score *= dist
			}
			if score > maxscore {
				maxscore = score
			}
		}
	}
	fmt.Println(maxscore)
}
