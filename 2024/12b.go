package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

//go:embed inputs/12.txt
var input string

func main() {
	G := grid.BytesFromStrings(exp.NonEmpty(strings.Split(input, "\n")))
	bounds := G.Bounds()
	u := make(algo.DisjointSets[image.Point])
	for p, c := range G.All {
		u.Find(p)
		for _, d := range []image.Point{{1, 0}, {0, 1}} {
			q := p.Add(d)
			if !q.In(bounds) {
				continue
			}
			if G.At(q) == c {
				u.Union(p, q)
			}
		}
	}

	sum := 0
	for _, sl := range u.Sets() {
		s := algo.SetFromSlice(sl)
		type edge struct {
			p, q image.Point
		}
		sides := make(algo.DisjointSets[edge])
		for p := range s {
			for _, d := range algo.Neigh4 {
				if q := p.Add(d); !s.Contains(q) {
					sides.Find(edge{p, q})
				}
			}
		}
		for e := range sides {
			for _, d := range algo.Neigh4 {
				f := edge{e.p.Add(d), e.q.Add(d)}
				if _, ok := sides[f]; ok {
					sides.Union(e, f)
				}
			}
		}

		sum += len(sides.Reps()) * len(s)
	}
	fmt.Println(sum)
}
