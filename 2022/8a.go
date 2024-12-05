package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2022
// Day 8, part a

func main() {
	g := exp.MustReadByteGrid("inputs/8.txt")

	bounds := g.Bounds()

	vc := 0
	for r := range g {
		for c, h := range g[r] {
			dvis := false
			for _, d := range algo.Neigh4 {
				vis := true
				p := image.Pt(c, r).Add(d)
				for p.In(bounds) {
					if g[p.Y][p.X] >= h {
						vis = false
						break
					}
					p = p.Add(d)
				}
				if vis {
					dvis = true
					break
				}
			}
			if dvis {
				vc++
			}
		}
	}
	fmt.Println(vc)
}
