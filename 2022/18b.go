package main

import (
	"errors"
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2022
// Day 18, part b

var neigh = []algo.Vec3[int]{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func main() {
	var min, max algo.Vec3[int]
	lava := make(algo.Set[algo.Vec3[int]])
	for i, line := range exp.MustReadLines("inputs/18.txt") {
		var v algo.Vec3[int]
		exp.Must(fmt.Sscanf(line, "%d,%d,%d", &v[0], &v[1], &v[2]))
		lava.Insert(v)
		if i == 0 {
			min, max = v, v
			continue
		}
		min[0] = algo.Min(min[0], v[0])
		min[1] = algo.Min(min[1], v[1])
		min[2] = algo.Min(min[2], v[2])
		max[0] = algo.Max(max[0], v[0])
		max[1] = algo.Max(max[1], v[1])
		max[2] = algo.Max(max[2], v[2])
	}

	intErr := errors.New("interior face")
	extErr := errors.New("exterior face")
	inside := make(algo.Set[algo.Vec3[int]])
	outside := make(algo.Set[algo.Vec3[int]])
	outside.Insert(min.Add(neigh[1]))
	outside.Insert(max.Add(neigh[0]))
	sa := 0
	for c := range lava {
		for _, n := range neigh {
			p := c.Add(n)
			if lava.Contains(p) {
				continue
			}
			if _, err := algo.FloodFill(p, func(x algo.Vec3[int], _ int) ([]algo.Vec3[int], error) {
				if inside.Contains(x) {
					return nil, intErr
				}
				if outside.Contains(x) {
					return nil, extErr
				}
				var next []algo.Vec3[int]
				for _, n := range neigh {
					if lava.Contains(x.Add(n)) {
						continue
					}
					next = append(next, x.Add(n))
				}
				return next, nil
			}); err != extErr {
				inside.Insert(p)
				continue
			}
			outside.Insert(p)
			sa++
		}
	}

	fmt.Println(sa)
}
