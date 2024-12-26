package main

import (
	"fmt"
	"slices"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2024
// Day 23, part a

const inputPath = "2024/inputs/23.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	sets := make(map[string]algo.Set[string])
	for _, line := range lines {
		u, v := exp.MustCut(line, "-")
		sets[u] = sets[u].Insert(v)
		sets[v] = sets[v].Insert(u)
	}
	triples := make(algo.Set[[3]string])
	for u, ua := range sets {
		for v := range ua {
			va := sets[v]
			for w := range va {
				if u == w || v == w {
					continue
				}
				if u[0] != 't' && v[0] != 't' && w[0] != 't' {
					continue
				}
				if !sets[u].Contains(w) {
					continue
				}
				trip := [3]string{u, v, w}
				slices.Sort(trip[:])
				triples.Insert(trip)
			}
		}
	}
	fmt.Println(len(triples))
}
