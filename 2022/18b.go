package main

import (
	"errors"
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
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
	var min algo.Vec3[int]
	lava := make(algo.Set[algo.Vec3[int]])
	for i, line := range exp.MustReadLines("inputs/18.txt") {
		var v algo.Vec3[int]
		exp.Must(fmt.Sscanf(line, "%d,%d,%d", &v[0], &v[1], &v[2]))
		lava.Insert(v)
		if i == 0 {
			min = v
			continue
		}
		min[0] = algo.Min(min[0], v[0])
		min[1] = algo.Min(min[1], v[1])
		min[2] = algo.Min(min[2], v[2])
	}

	ext := errors.New("exterior face")
	outside := algo.Set[algo.Vec3[int]]{min.Add(neigh[1]): {}}
	sa := 0
	for c := range lava {
		for _, n := range neigh {
			if lava.Contains(c.Add(n)) {
				continue
			}
			if _, err := algo.FloodFill(c.Add(n), func(x algo.Vec3[int], _ int) ([]algo.Vec3[int], error) {
				if outside.Contains(x) {
					return nil, ext
				}
				var next []algo.Vec3[int]
				for _, n := range neigh {
					if lava.Contains(x.Add(n)) {
						continue
					}
					next = append(next, x.Add(n))
				}
				return next, nil
			}); err != ext {
				continue
			}
			outside.Insert(c.Add(n))
			sa++
		}
	}

	fmt.Println(sa)
}
