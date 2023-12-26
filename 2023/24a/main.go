package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 24, part a

const inputPath = "2023/inputs/24.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	type stone struct {
		p, v algo.Vec3[int]
	}

	stones := algo.Map(lines, func(line string) (s stone) {
		exp.MustSscanf(line, "%d, %d, %d @ %d, %d, %d", &s.p[0], &s.p[1], &s.p[2], &s.v[0], &s.v[1], &s.v[2])
		return s
	})

	rng := algo.Range[int]{Min: 200000000000000, Max: 400000000000000}

	count := 0

	for i, s1 := range stones[:len(stones)-1] {
		for _, s2 := range stones[i+1:] {
			// x1(t) = x1(0) + vx1 * t; x2(t) = x2(0) + vx2 * t;
			// y1(t) = y1(0) + vy1 * t; y2(t) = y2(0) + vy2 * t;
			// So looking for x1(t) = x2(t), y1(t) = y2(t),
			//  with t > 0, and
			//     x1(0) + vx1 * t = x2(0) + vx2 * t
			// =>    x1(0) - x2(0) = (vx2 - vx1) * t;
			// sim.  y1(0) - y2(0) = (vy2 - vy1) * t.

		}
	}

	fmt.Println(count)
}
