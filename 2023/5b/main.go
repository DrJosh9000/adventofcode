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
// Day 5, part b

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
		rng algo.Range[int]
	}

	minLoc := math.MaxInt
	var q []catIdx
	for i := 0; i < len(seeds); i += 2 {
		q = append(q, catIdx{cat: "seed", rng: algo.Range[int]{Min: seeds[i], Max: seeds[i] + seeds[i+1] - 1}})
	}

	for len(q) > 0 {
		ci := q[0]
		q = q[1:]
		//log.Printf("considering %v (len(q) = %d)", ci, len(q))

		if ci.rng.IsEmpty() {
			continue
		}

		if ci.cat == "location" {
			minLoc = min(minLoc, ci.rng.Min)
			continue
		}

		for p, rms := range maps {
			if p.src != ci.cat {
				continue
			}

			intersectAny := false
			for _, r := range rms {
				rx := algo.Range[int]{Min: r.srcStart, Max: r.srcStart + r.length - 1}
				o := rx.Intersection(ci.rng)
				if o.IsEmpty() {
					continue
				}

				intersectAny = true
				for _, rem := range subtract(ci.rng, rx) {
					if rem.IsEmpty() {
						continue
					}
					nci := catIdx{cat: ci.cat, rng: rem}
					//log.Printf("pushing %v (subtract)", nci)
					q = append(q, nci)
				}

				delta := r.dstStart - r.srcStart
				o.Min += delta
				o.Max += delta
				nci := catIdx{cat: p.dst, rng: o}
				//log.Printf("pushing %v (intersect with %v)", nci, rx)
				q = append(q, nci)
				break
			}

			if !intersectAny {
				nci := catIdx{cat: p.dst, rng: ci.rng}
				//log.Printf("pushing %v (no intersect)", nci)
				q = append(q, nci)
			}
		}
	}

	fmt.Println(minLoc)
}

func subtract(a, b algo.Range[int]) []algo.Range[int] {
	if b.Min <= a.Min && b.Max >= a.Max {
		return nil
	}
	if b.Min >= a.Min && b.Max <= a.Max {
		return []algo.Range[int]{
			{Min: a.Min, Max: b.Min},
			{Min: b.Max + 1, Max: a.Max},
		}
	}
	if b.Min >= a.Min {
		a.Max = min(a.Max, b.Min-1)
	} else {
		a.Min = max(a.Min, b.Max+1)
	}
	return []algo.Range[int]{a}
}
