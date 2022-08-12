package main

import (
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/DrJosh9000/exp"
)

type nanobot struct {
	x, y, z, r int
}

func (b nanobot) norm() int {
	return abs(b.x) + abs(b.y) + abs(b.z)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(v, w nanobot) int {
	return abs(v.x - w.x) + abs(v.y - w.y) + abs(v.z - w.z)
}

// overlaps reports if v and w intersect in any way
func overlaps(v, w nanobot) bool {
	return dist(v, w) < v.r + w.r
}

// contains reports if v contains w
func contains(v, w nanobot) bool {
	return dist(v, w) + w.r <= v.r
}

// union returns a sphere containing both v and w
func union(v, w nanobot) nanobot {
	if contains(v, w) {
		return v
	}
	if contains(w, v) {
		return w
	}
	minx := min(v.x - v.r, w.x - w.r)
	maxx := max(v.x + v.r, w.x + w.r)
	miny := min(v.y - v.r, w.y - w.r)
	maxy := max(v.y + v.r, w.y + w.r)
	minz := min(v.z - v.r, w.z - w.r)
	maxz := max(v.z + v.r, w.z + w.r)
	z := nanobot{
		x: (minx + maxx) / 2,
		y: (miny + maxy) / 2,
		z: (minz + maxz) / 2,
		r: (max(max(maxx - minx, maxy - miny), maxz - minz) + 1) / 2,
	}
	return z
}
var bots []nanobot

func search(bounds nanobot) (count, norm int) {
	r := bounds.r / 2
	next := []nanobot{
		{bounds.x - r, bounds.y, bounds.z, r},
		{bounds.x + r, bounds.y, bounds.z, r},
		{bounds.x, bounds.y - r, bounds.z, r},
		{bounds.x, bounds.y + r, bounds.z, r},
		{bounds.x, bounds.y, bounds.z - r, r},
		{bounds.x, bounds.y, bounds.z + r, r},
	}
	var bestbounds []nanobot
	var bestcount int
	for _, bound := range next {
		count := 0
		for _, b := range bots {
			if overlaps(b, bound) {
				count++
			}
		}
		if len(bestbounds) == 0 || count > bestcount {
			bestbounds = []nanobot{bound}
			bestcount = count
			continue
		}
		if count == bestcount {
			bestbounds = append(bestbounds, bound)
		}
	}
	
	if r == 0 {
		bestnorm := math.MaxInt 
		for _, b := range bestbounds {
			if n := b.norm(); n < bestnorm {
				bestnorm = n
			}	
		}
		return bestcount, bestnorm
	}
	
	var wg sync.WaitGroup
	N := len(bestbounds)
	wg.Add(N)
	counts, norms := make([]int, N), make([]int, N)
	s := func(i int) {
		counts[i], norms[i] = search(bestbounds[i])
		wg.Done()
	}
	for i := range bestbounds {
		if r > 1000 {
			go s(i)
		} else {
			s(i)
		}
	}
	wg.Wait()
	
	bestcount = -1
	bestnorm := math.MaxInt
	for i := range counts {
		if counts[i] > bestcount || (counts[i] == bestcount && norms[i] < bestnorm) {
			bestcount = counts[i]
			bestnorm = norms[i]
		}
	}
	return bestcount, bestnorm
}

func main() {
	var bounds nanobot
	exp.MustForEachLineIn("inputs/23.txt", func(line string) {
		var bot nanobot
		if _, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &bot.x, &bot.y, &bot.z, &bot.r); err != nil {
			log.Fatalf("Couldn't scan line: %v", err)
		}
		if len(bots) == 0 {
			bounds = bot
		} else {
			bounds = union(bounds, bot)
		}
		bots = append(bots, bot)
	})
	
	count, norm := search(bounds)
	fmt.Println("Number of bots overlapping:", count)
	fmt.Println("Distance from origin:", norm)
}