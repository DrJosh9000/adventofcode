package main

import (
	"errors"
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 22, part b

func main() {
	// The assumptions given in the part 2 description seem to hold:
	// There is indeed only one empty node, that can handle small but not large
	// amounts of data.
	// So we only need to track where the hole is, and where the goal data is.
	// (A more general solution is possible, it merely isn't necessary.)
	type state struct {
		hole, goal image.Point
	}
	var start state
	lim := 0

	type node struct {
		size, used, avail, usep int
	}
	nodes := make(map[image.Point]node)
	for _, line := range exp.MustReadLines("inputs/22.txt") {
		if !strings.HasPrefix(line, "/dev") {
			continue
		}
		var n node
		var p image.Point
		exp.Must(fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT %d%%", &p.X, &p.Y, &n.size, &n.used, &n.avail, &n.usep))
		nodes[p] = n

		if n.used == 0 {
			start.hole = p
			lim = n.avail
		}
		if p.Y == 0 && p.X > start.goal.X {
			start.goal = p
		}
	}

	algo.FloodFill(start, func(s state, d int) ([]state, error) {
		if s.goal == (image.Point{}) {
			fmt.Println(d)
			return nil, errors.New("all done")
		}
		var out []state
		for _, Δ := range algo.Neigh4 {
			p := s.hole.Add(Δ)
			if n, ok := nodes[p]; !ok || n.used > lim {
				// doesn't exist or too big
				continue
			}
			if p == s.goal {
				out = append(out, state{hole: s.goal, goal: s.hole})
			} else {
				out = append(out, state{hole: p, goal: s.goal})
			}
		}
		return out, nil
	})
}
