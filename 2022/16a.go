package main

import (
	"fmt"
	"image"
	"log"
	"sort"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2022
// Day 16, part a

type valve struct {
	name    string
	rate    int
	tunsrc  []string
	tunnels []int
}

func main() {
	input := exp.MustReadLines("inputs/16.txt")
	var valves []*valve
	ngood := 0
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
		if v.rate > 0 {
			ngood++
		}
	}

	// Put all the good ones first
	sort.Slice(valves, func(i, j int) bool {
		return valves[i].rate > valves[j].rate
	})

	// Convert tunnels to indices
	label := make(map[string]int)
	for i, v := range valves {
		label[v.name] = i
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
	allsrc := []int{aa}
	for i := 0; i < ngood; i++ {
		allsrc = append(allsrc, i)
	}

	// Find shortest paths between good valves
	dist := make(grid.Sparse[int])
	for _, i := range allsrc {
		algo.FloodFill(i, func(j, d int) ([]int, error) {
			if j < ngood {
				p := image.Pt(i, j)
				if old, exists := dist[p]; !exists || d < old {
					dist[p] = d
				}
			}
			return valves[j].tunnels, nil
		})
	}

	const timelimit = 30

	type state struct {
		pos  int
		open uint
	}
	tab := []map[state]int{
		{{pos: aa}: 0},
	}
	for len(tab) <= timelimit {
		tab = append(tab, make(map[state]int))
	}

	maxrel := 0
	for t, m := range tab[:timelimit] {
		// Starting with every state at time t...
		for s, r := range m {
			// consider opening the valve at s.pos, if not open
			if s.open&(1<<s.pos) == 0 {
				// the valve is considered open from the next timestep
				ns := s
				ns.open |= 1 << s.pos
				nr := r + (timelimit-t-1)*valves[s.pos].rate
				if or, exist := tab[t+1][ns]; !exist || or < nr {
					tab[t+1][ns] = nr
					if maxrel < nr {
						maxrel = nr
					}
				}
			}
			// consider moving to another closed valve j
			for j := 0; j < ngood; j++ {
				if s.pos == j {
					continue
				}
				if s.open&(1<<j) != 0 {
					continue
				}
				dt := dist[image.Pt(s.pos, j)]
				if t+dt > timelimit {
					continue
				}
				ns := s
				ns.pos = j
				if or, exist := tab[t+dt][ns]; !exist || or < r {
					tab[t+dt][ns] = r
				}
			}
		}
	}

	fmt.Println(maxrel)
}
