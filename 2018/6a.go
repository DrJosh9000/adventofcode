package main

import (
	"fmt"
	"image"
	"log"
	"math"

	"drjosh.dev/exp"
)

var step = []image.Point{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func l1(p image.Point) int {
	return abs(p.X) + abs(p.Y)
}

func main() {
	var pts []image.Point
	var bounds image.Rectangle
	bounds.Min = image.Pt(math.MaxInt, math.MaxInt)
	bounds.Max = image.Pt(math.MinInt, math.MinInt)
	exp.MustForEachLineIn("inputs/6.txt", func(line string) {
		var p image.Point
		if _, err := fmt.Sscanf(line, "%d, %d", &p.X, &p.Y); err != nil {
			log.Fatalf("Couldn't parse line %q: %v", line, err)
		}
		pts = append(pts, p)
		if p.X < bounds.Min.X {
			bounds.Min.X = p.X
		}
		if p.X > bounds.Max.X {
			bounds.Max.X = p.X
		}
		if p.Y < bounds.Min.Y {
			bounds.Min.Y = p.Y
		}
		if p.Y > bounds.Max.Y {
			bounds.Max.Y = p.Y
		}
	})

	nearest := func(p image.Point) int {
		min := math.MaxInt
		best := -1
		for i, q := range pts {
			t := l1(p.Sub(q))
			if t == min {
				best = -1
				continue
			}
			if t < min {
				min = t
				best = i
			}
		}
		return best
	}

	// We want a L1 discrete Voronoi diagram ...
	// Here's a simple way. Test each point in the bounds against each
	// point in the list to see which is closest.
	sizes := make([]int, len(pts))
	var p image.Point
	for p.X = bounds.Min.X; p.X <= bounds.Max.X; p.X++ {
		for p.Y = bounds.Min.Y; p.Y <= bounds.Max.Y; p.Y++ {
			best := nearest(p)
			if best != -1 {
				sizes[best]++
			}
		}
	}

	// Consider the boundary to find points within infinite cells.
	cand := make(map[int]struct{})
	for i := range pts {
		cand[i] = struct{}{}
	}
	for p.X = bounds.Min.X; p.X <= bounds.Max.X; p.X++ {
		p.Y = bounds.Min.Y
		if t := nearest(p); t != -1 {
			delete(cand, t)
		}
		p.Y = bounds.Max.Y
		if t := nearest(p); t != -1 {
			delete(cand, t)
		}
	}
	for p.Y = bounds.Min.Y; p.Y <= bounds.Max.Y; p.Y++ {
		p.X = bounds.Min.X
		if t := nearest(p); t != -1 {
			delete(cand, t)
		}
		p.X = bounds.Max.X
		if t := nearest(p); t != -1 {
			delete(cand, t)
		}
	}

	max := math.MinInt
	for i := range cand {
		if sizes[i] > max {
			max = sizes[i]
		}
	}
	fmt.Println(max)
}
