package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 5, part a

func main() {
	var seeds []int

	type maprange struct {
		dstStart, srcStart, length int
	}

	type catPair struct {
		src, dst string
	}

	var currentCat catPair

	maps := make(map[catPair][]maprange)

	for _, line := range exp.MustReadLines("2023/inputs/5.txt") {
		switch {
		case strings.HasPrefix(line, "seeds:"):
			seeds = algo.MustMap(strings.Fields(strings.TrimPrefix(line, "seeds: ")), strconv.Atoi)

		case strings.HasSuffix(line, "map:"):
			var ok bool
			currentCat.src, currentCat.dst, ok = strings.Cut(strings.TrimSuffix(line, " map:"), "-to-")
			if !ok {
				log.Fatalf("No -to-: %q", line)
			}

		case line == "":
			// nothing

		default:
			var rng maprange
			exp.Must(fmt.Sscanf(line, "%d %d %d", &rng.dstStart, &rng.srcStart, &rng.length))
			maps[currentCat] = append(maps[currentCat], rng)
		}
	}

	type catIdx struct {
		cat string
		idx int
	}

	minLoc := math.MaxInt
	for _, seed := range seeds {
		q := []catIdx{{cat: "seed", idx: seed}}
		for len(q) > 0 {
			ci := q[0]
			q = q[1:]

			for p, rs := range maps {
				if p.src != ci.cat {
					continue
				}
				dst := ci.idx
				for _, r := range rs {
					if ci.idx >= r.srcStart && ci.idx < r.srcStart+r.length {
						dst = r.dstStart + ci.idx - r.srcStart
						break
					}
				}
				if p.dst == "location" {
					minLoc = min(minLoc, dst)
				}
				q = append(q, catIdx{cat: p.dst, idx: dst})
			}
		}
	}

	fmt.Println(minLoc)
}
