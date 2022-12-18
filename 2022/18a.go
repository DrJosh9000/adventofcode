package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 18, part a

var neigh = []algo.Vec3[int]{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func main() {
	lava := make(algo.Set[algo.Vec3[int]])
	for _, line := range exp.MustReadLines("inputs/18.txt") {
		var v algo.Vec3[int]
		exp.Must(fmt.Sscanf(line, "%d,%d,%d", &v[0], &v[1], &v[2]))
		lava.Insert(v)
	}

	sa := 0
	for c := range lava {
		for _, n := range neigh {
			if !lava.Contains(c.Add(n)) {
				sa++
			}
		}
	}

	fmt.Println(sa)
}
