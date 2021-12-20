package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/19.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var scans [][]vec
	var probes []vec
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "---") {
			if len(probes) > 0 {
				scans = append(scans, probes)
				probes = nil
			}
			continue
		}
		var probe vec
		if _, err := fmt.Sscanf(line, "%d,%d,%d", &probe.x, &probe.y, &probe.z); err != nil {
			log.Fatalf("Couldn't sscanf: %v", err)
		}
		probes = append(probes, probe)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}
	if len(probes) > 0 {
		scans = append(scans, probes)
	}

	// Let scan 0 define the reference frame for the whole puzzle. Seed the
	// universe with scan 0.
	universe := make(set, len(scans[0]))
	universe.subsume(scans[0])
	scanners := make(set, len(scans))
	scanners[vec{}] = struct{}{}

	// For all the other scans...
	remaining := make(map[int]struct{})
	for i := range scans[1:] {
		remaining[i+1] = struct{}{}
	}
remainingLoop:
	for len(remaining) > 0 {
		for i := range remaining {
			//log.Printf("Trying scan %d", i)
			scan := scans[i]
			for _, r := range rotations {
				rotated := make([]vec, 0, len(scan))
				for _, probe := range scan {
					rotated = append(rotated, probe.rotate(r))
				}
				// Try each known as the origin for the scan.
				for known := range universe {
					// Try each probe as the origin for the scan.
					for _, probe := range rotated {
						origin := probe.sub(known)
						piece := make([]vec, 0, len(scan))
						for _, p := range rotated {
							piece = append(piece, p.sub(origin))
						}
						comm := universe.common(piece)
						//log.Printf("Scan %d in rotation %d with origin %v had %d in common", i, r, origin, comm)
						if comm >= 12 {
							//log.Printf("Matched scan %d", i)
							universe.subsume(piece)
							delete(remaining, i)
							scanners[vec{}.sub(origin)] = struct{}{}
							continue remainingLoop
						}
					}

				}
			}
		}
		log.Fatal("Ran out of scans to try")
	}

	fmt.Printf("Size of universe: %d\n", len(universe))
	max := 0
	for p := range scanners {
		for q := range scanners {
			if d := p.dist(q); d > max {
				max = d
			}
		}
	}
	fmt.Printf("Biggest Manhattan distance between scanners: %v\n", max)
}

type set map[vec]struct{}

func (s set) subsume(t []vec) {
	for _, v := range t {
		s[v] = struct{}{}
	}
}

func (s set) common(scan []vec) int {
	c := 0
	for _, v := range scan {
		if _, found := s[v]; found {
			c++
		}
	}
	return c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type vec struct{ x, y, z int }

func (v vec) dist(w vec) int {
	return abs(v.x-w.x) + abs(v.y-w.y) + abs(v.z-w.z)
}

func (v vec) sub(w vec) vec {
	return vec{v.x - w.x, v.y - w.y, v.z - w.z}
}

func (v vec) a(n int) vec {
	switch n {
	case 0:
		return v
	case 1:
		return vec{v.y, v.z, v.x}
	case 2:
		return vec{v.z, v.x, v.y}
	}
	panic("n out of range")
}

func (v vec) b(n int) vec {
	switch n {
	case 0:
		return v
	case 1:
		return vec{-v.x, -v.y, v.z}
	case 2:
		return vec{-v.x, v.y, -v.z}
	case 3:
		return vec{v.x, -v.y, -v.z}
	}
	panic("n out of range")
}

func (v vec) c(n int) vec {
	switch n {
	case 0:
		return v
	case 1:
		return vec{-v.z, -v.y, -v.x}
	}
	panic("n out of range")
}

func (v vec) rotate(r vec) vec {
	return v.a(r.x).b(r.y).c(r.z)
}

var rotations = []vec{
	{0, 0, 0},
	{0, 0, 1},
	{0, 1, 0},
	{0, 1, 1},
	{0, 2, 0},
	{0, 2, 1},
	{0, 3, 0},
	{0, 3, 1},
	{1, 0, 0},
	{1, 0, 1},
	{1, 1, 0},
	{1, 1, 1},
	{1, 2, 0},
	{1, 2, 1},
	{1, 3, 0},
	{1, 3, 1},
	{2, 0, 0},
	{2, 0, 1},
	{2, 1, 0},
	{2, 1, 1},
	{2, 2, 0},
	{2, 2, 1},
	{2, 3, 0},
	{2, 3, 1},
}
