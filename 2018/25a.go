package main

import (
	"fmt"
	"log"

	"drjosh.dev/exp"
)

type set[T comparable] map[T]struct{}

func (s set[T]) add(t T) { s[t] = struct{}{} }

type point struct {
	x, y, z, t int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(p, q point) int {
	return abs(p.x-q.x) + abs(p.y-q.y) + abs(p.z-q.z) + abs(p.t-q.t)
}

func main() {
	var pts []point
	var up []int
	exp.MustForEachLineIn("inputs/25.txt", func(line string) {
		var p point
		if _, err := fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.t); err != nil {
			log.Fatalf("Couldn't scan line %q: %v", line, err)
		}
		con := len(up)
		up = append(up, con)
		for j, q := range pts {
			if dist(p, q) <= 3 {
				u := j
				for u != up[u] {
					u = up[u]
				}
				up[u] = con
			}
		}
		pts = append(pts, p)
	})

	cons := make(set[int])
	for i := range pts {
		for up[i] != up[up[i]] {
			up[i] = up[up[i]]
		}
		cons.add(up[i])
	}
	fmt.Println(len(cons))
}
