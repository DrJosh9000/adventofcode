package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2016
// Day 11, part b

type state struct {
	rtgs     [7]byte
	chips    [7]byte
	elevator byte
}

func (s state) valid() bool {
	// every chip...
	for c, cf := range s.chips {
		// ...is safe if on the same floor as its rtg...
		if cf == s.rtgs[c] {
			continue
		}
		// ...but is unsafe on any floor with any other rtg.
		for r, rf := range s.rtgs {
			if r != c && rf == cf {
				return false
			}
		}
		// here c must be on a floor with no rtg.
	}
	return true
}

var ordinal = map[string]byte{
	"first":  0,
	"second": 1,
	"third":  2,
	"fourth": 3,
}

func main() {
	start := state{
		elevator: 0,
	}
	elements := make(map[string]int)
	element := func(e string) int {
		n, exists := elements[e]
		if !exists {
			n = len(elements)
			elements[e] = n
		}
		return n
	}
	for _, line := range exp.MustReadLines("inputs/11.txt") {
		if strings.HasSuffix(line, "nothing relevant.") {
			continue
		}
		parts := strings.Split(line, " floor contains ")
		floor := ordinal[strings.TrimPrefix(parts[0], "The ")]
		parts[1] = strings.TrimSuffix(parts[1], ".")
		for _, item := range strings.Split(parts[1], ", ") {
			// one line has "a foo and a bar"
			// one line has an Oxford comma
			// maybe I should hand-code the start state...
			for _, item := range strings.Split(item, " and ") {
				if item == "" {
					continue
				}
				//log.Printf("item: %q", item)
				item = strings.TrimPrefix(item, "and ")
				item = strings.TrimPrefix(item, "a ")
				switch {
				case strings.HasSuffix(item, "generator"):
					e := strings.TrimSuffix(item, " generator")
					start.rtgs[element(e)] = floor
				case strings.HasSuffix(item, "microchip"):
					e := strings.TrimSuffix(item, "-compatible microchip")
					start.chips[element(e)] = floor
				default:
					log.Fatalf("Unknown item: %q", item)
				}
			}
		}
	}
	// These are on the first floor; zero value suffices.
	elements["elerium"] = 5
	elements["dilithium"] = 6
	goal := state{
		elevator: 3,
	}
	for _, n := range elements {
		goal.rtgs[n] = 3
		goal.chips[n] = 3
	}

	algo.FloodFill(start, func(s state, d int) ([]state, error) {
		if s == goal {
			fmt.Println(d)
			return nil, errors.New("all done")
		}
		var next []state
		var nfs []byte
		if s.elevator > 0 {
			nfs = append(nfs, s.elevator-1)
		}
		if s.elevator < 3 {
			nfs = append(nfs, s.elevator+1)
		}
		for _, nf := range nfs {
			n := s
			n.elevator = nf
			// consider moving one rtg from f to nf
			for r, f := range s.rtgs {
				if f != s.elevator {
					continue
				}
				n.rtgs[r] = nf
				if n.valid() {
					next = append(next, n)
				}
				// consider moving two rtgs from f to nf
				for r2, f2 := range s.rtgs {
					if r2 == r || f2 != s.elevator {
						continue
					}
					n.rtgs[r2] = nf
					if n.valid() {
						next = append(next, n)
					}
					n.rtgs[r2] = s.elevator
				}
				// consider moving an rtg and a chip from this floor
				for c, f2 := range s.chips {
					if f2 != s.elevator {
						continue
					}
					n.chips[c] = nf
					if n.valid() {
						next = append(next, n)
					}
					n.chips[c] = s.elevator
				}
				n.rtgs[r] = s.elevator
			}
			// consider moving one chip from this floor
			for c, f := range s.chips {
				if f != s.elevator {
					continue
				}
				n.chips[c] = nf
				if n.valid() {
					next = append(next, n)
				}
				// consider moving two chips from this floor
				for c2, f2 := range s.chips {
					if f2 != s.elevator {
						continue
					}
					n.chips[c2] = nf
					if n.valid() {
						next = append(next, n)
					}
					n.chips[c2] = s.elevator
				}
				n.chips[c] = s.elevator
			}
		}
		return next, nil
	})
}
