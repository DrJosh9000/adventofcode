package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 16, part b

type valve struct {
	name    string
	rate    int16
	tunsrc  []string
	tunnels []int
}

type state uint

// 6 + 6 + 15 = 27

func (s state) p1() int {
	return int(s) >> 21
}

func (s state) p2() int {
	return (int(s) >> 15) & 0x3f
}

func (s state) open() uint {
	return uint(s) & 0x7fff
}

func pack(p1, p2 int, open uint) state {
	return state(p1)<<21 | state(p2)<<15 | state(open)
}

func main() {
	cpuf := exp.Must(os.Create("cpu-16b-fast.prof"))
	defer cpuf.Close()
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

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

	// Put all the good ones first
	sort.Slice(valves, func(i, j int) bool {
		return valves[i].rate > valves[j].rate
	})

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

	tab := make([]int16, 1<<27)
	for i := range tab {
		tab[i] = -1
	}
	tab[pack(aa, aa, 0)] = 0

	var maxrel int16
	for t := int16(0); t < timelimit; t++ {
		log.Printf("Commencing timestep %d - considering %d states\n", t, len(tab))

		t2 := make([]int16, 1<<27)
		for i := range t2 {
			t2[i] = -1
		}

		upd := func(s state, r int16) {
			if t2[s] < r {
				t2[s] = r
			}
		}

		for si, r := range tab {
			if r < 0 {
				continue
			}
			s := state(si)
			v1 := valves[s.p1()]
			v2 := valves[s.p2()]

			// Both move
			for _, j := range v1.tunnels {
				for _, k := range v2.tunnels {
					upd(pack(j, k, s.open()), r)
				}
			}

			if v1.rate > 0 && s.open()&(1<<s.p1()) == 0 {
				// p1 has the option of opening a valve, while p2 moves
				nr := r + (timelimit-t-1)*v1.rate
				if maxrel < nr {
					maxrel = nr
				}
				if open := s.open() | 1<<s.p1(); open != allopen {
					for _, j := range v2.tunnels {
						upd(pack(s.p1(), j, open), nr)
					}
				}
			}

			if v2.rate > 0 && s.open()&(1<<s.p2()) == 0 {
				// p2 has the option of opening a valve, while p1 moves.
				nr := r + (timelimit-t-1)*v2.rate
				if maxrel < nr {
					maxrel = nr
				}

				if open := s.open() | 1<<s.p2(); open != allopen {
					for _, j := range v1.tunnels {
						upd(pack(j, s.p2(), open), nr)
					}
				}
			}

			// also consider both opening, if that's possible
			// ... make sure they are at separate valves!
			if v1.rate > 0 && s.open()&(1<<s.p1()) == 0 && v2.rate > 0 && s.open()&(1<<s.p2()) == 0 && s.p1() != s.p2() {
				nr := r + (timelimit-t-1)*(v1.rate+v2.rate)
				if maxrel < nr {
					maxrel = nr
				}
				if open := s.open() | 1<<s.p1() | 1<<s.p2(); open != allopen {
					upd(pack(s.p1(), s.p2(), open), nr)
				}
			}

		}

		tab = t2
	}

	fmt.Println(maxrel)
}
