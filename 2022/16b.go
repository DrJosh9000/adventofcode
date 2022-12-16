package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 16, part b

type valve struct {
	name    string
	rate    int
	tunsrc  []string
	tunnels []int
}

func main() {
	input := exp.MustReadLines("inputs/16.txt")
	var valves []*valve
	for _, line := range input {
		var v valve
		a, b, ok := strings.Cut(line, ";")
		if !ok {
			log.Fatalf("Line missing cut string: %q", line)
		}
		exp.Must(fmt.Sscanf(a, "Valve %s has flow rate=%d", &v.name, &v.rate))
		b = strings.TrimPrefix(b, " tunnels lead to valves ")
		b = strings.TrimPrefix(b, " tunnel leads to valve ")
		v.tunsrc = strings.Split(b, ", ")
		valves = append(valves, &v)
	}

	var allopen uint

	// Convert tunnels to indices
	label := make(map[string]int)
	for i, v := range valves {
		label[v.name] = i
		if v.rate > 0 {
			allopen |= 1 << i
		}
	}

	for _, v := range valves {
		v.tunnels = make([]int, 0, len(v.tunsrc))
		for _, t := range v.tunsrc {
			if _, nope := label[t]; !nope {
				fmt.Printf("Hmm, %q is not a label\n", t)
			}
			v.tunnels = append(v.tunnels, label[t])
		}
	}

	// Where do we start again?
	aa := label["AA"]

	const timelimit = 26

	type state struct {
		p1, p2 int
		open   uint
	}

	tab := map[state]int{
		{p1: aa, p2: aa}: 0,
	}

	maxrel := 0
	for t := 0; t < timelimit; t++ {
		log.Printf("Commencing timestep %d - considering %d states\n", t, len(tab))

		t2 := make(map[state]int)

		upd := func(s state, r int) {
			if or, exist := t2[s]; !exist || or < r {
				t2[s] = r
			}
		}

		for s, r := range tab {
			v1 := valves[s.p1]
			v2 := valves[s.p2]

			// Both move
			ns := s
			for _, ns.p1 = range v1.tunnels {
				for _, ns.p2 = range v2.tunnels {
					upd(ns, r)
				}
			}

			if v1.rate > 0 && s.open&(1<<s.p1) == 0 {
				// p1 has the option of opening a valve, while p2 moves
				nr := r + (timelimit-t-1)*v1.rate
				if maxrel < nr {
					maxrel = nr
				}
				ns := s
				ns.open |= 1 << s.p1
				if ns.open != allopen {
					for _, ns.p2 = range v2.tunnels {
						upd(ns, nr)
					}
				}
			}

			if v2.rate > 0 && s.open&(1<<s.p2) == 0 {
				// p2 has the option of opening a valve, while p1 moves.
				nr := r + (timelimit-t-1)*v2.rate
				if maxrel < nr {
					maxrel = nr
				}
				ns := s
				ns.open |= 1 << s.p2
				if ns.open != allopen {
					for _, ns.p1 = range v1.tunnels {
						upd(ns, nr)
					}
				}
			}

			// also consider both opening, if that's possible
			// ... make sure they are at separate valves!
			if v1.rate > 0 && s.open&(1<<s.p1) == 0 && v2.rate > 0 && s.open&(1<<s.p2) == 0 && s.p1 != s.p2 {
				nr := r + (timelimit-t-1)*(v1.rate+v2.rate)
				if maxrel < nr {
					maxrel = nr
				}
				ns := s
				ns.open |= 1 << s.p1
				ns.open |= 1 << s.p2
				if ns.open != allopen {
					upd(ns, nr)
				}
			}

		}

		tab = t2
	}

	fmt.Println(maxrel)
}
