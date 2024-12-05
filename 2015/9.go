package main

import (
	"fmt"
	"math"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2015
// Day 9

// Travelling Salesman let's goooooo

func main() {
	label := make(map[string]int)
	input := exp.MustReadLines("inputs/9.txt")
	for _, line := range input {
		fs := strings.Fields(line)
		if _, l := label[fs[0]]; !l {
			label[fs[0]] = len(label)
		}
		if _, l := label[fs[2]]; !l {
			label[fs[2]] = len(label)
		}
	}

	N := len(label)

	dist := grid.Make[int](N, N)
	for _, line := range input {
		var s, t string
		var d int
		exp.Must(fmt.Sscanf(line, "%s to %s = %d", &s, &t, &d))
		x, y := label[s], label[t]
		if x == y {
			fmt.Printf("that's weird: %q and %q label the same node %d\n", s, t, x)
		}
		dist[x][y] = d
		dist[y][x] = d
	}

	fmt.Printf("distance matrix:\n%v\n", dist)

	ord := make([]int, N)
	for i := range ord {
		ord[i] = i
	}

	mind, maxd := math.MaxInt, math.MinInt
	for {
		d := 0
		p := ord[0]
		for _, q := range ord[1:] {
			d += dist[p][q]
			p = q
		}
		if d < mind {
			mind = d
		}
		if d > maxd {
			maxd = d
		}
		if !algo.NextPermutation(ord) {
			break
		}
	}
	fmt.Println(mind, maxd)
}
