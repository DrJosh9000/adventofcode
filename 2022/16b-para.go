package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/DrJosh9000/exp"
	"golang.org/x/exp/constraints"
)

// Advent of Code 2022
// Day 16, part b

const timelimit = 26

type valve struct {
	name    string
	rate    int
	tunsrc  []string
	tunnels []int
}

func mergeMax[K comparable, V constraints.Ordered](m ...map[K]V) map[K]V {
	mt := make(map[K]V)
	for _, m := range m {
		for k, v := range m {
			if w, ex := mt[k]; !ex || w < v {
				mt[k] = v
			}
		}
	}
	return mt
}

type state struct {
	p1, p2 int
	open   uint
}

type staterel struct {
	s state
	r int
}

func merge(outs ...[]staterel) map[state]int {
	s := 0
	for _, o := range outs {
		s += len(o)
	}
	mt := make(map[state]int, s)
	for _, o := range outs {
		for _, sr := range o {
			if w, ex := mt[sr.s]; !ex || w < sr.r {
				mt[sr.s] = sr.r
			}
		}
	}
	return mt
}

type info struct {
	maxrel  int
	valves  []*valve
	allopen uint
}

func (i *info) worker(t int, input []staterel) []staterel {
	out := make([]staterel, 0, 10*len(input))
	for _, sr := range input {
		s, r := sr.s, sr.r

		v1 := i.valves[s.p1]
		v2 := i.valves[s.p2]

		// Both move
		ns := s
		for _, ns.p1 = range v1.tunnels {
			for _, ns.p2 = range v2.tunnels {
				out = append(out, staterel{ns, r})
			}
		}

		if v1.rate > 0 && s.open&(1<<s.p1) == 0 {
			// p1 has the option of opening a valve, while p2 moves
			nr := r + (timelimit-t-1)*v1.rate
			if i.maxrel < nr {
				i.maxrel = nr
			}
			ns := s
			ns.open |= 1 << s.p1
			if ns.open != i.allopen {
				for _, ns.p2 = range v2.tunnels {
					out = append(out, staterel{ns, nr})
				}
			}
		}

		if v2.rate > 0 && s.open&(1<<s.p2) == 0 {
			// p2 has the option of opening a valve, while p1 moves.
			nr := r + (timelimit-t-1)*v2.rate
			if i.maxrel < nr {
				i.maxrel = nr
			}
			ns := s
			ns.open |= 1 << s.p2
			if ns.open != i.allopen {
				for _, ns.p1 = range v1.tunnels {
					out = append(out, staterel{ns, nr})
				}
			}
		}

		// also consider both opening, if that's possible
		// ... make sure they are at separate valves!
		if v1.rate > 0 && s.open&(1<<s.p1) == 0 && v2.rate > 0 && s.open&(1<<s.p2) == 0 && s.p1 != s.p2 {
			nr := r + (timelimit-t-1)*(v1.rate+v2.rate)
			if i.maxrel < nr {
				i.maxrel = nr
			}
			ns := s
			ns.open |= 1 << s.p1
			ns.open |= 1 << s.p2
			if ns.open != i.allopen {
				out = append(out, staterel{ns, nr})
			}
		}
	}
	return out
}

func main() {
	f := exp.Must(os.Create("cpu.pprof"))
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	go func() {
		<-sigch
		pprof.StopCPUProfile()
		f.Close()
		os.Exit(1)
	}()

	input := exp.MustReadLines("inputs/16.txt")
	var inf info

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
		inf.valves = append(inf.valves, &v)
	}

	// Convert tunnels to indices
	label := make(map[string]int)
	for i, v := range inf.valves {
		label[v.name] = i
		if v.rate > 0 {
			inf.allopen |= 1 << i
		}
	}

	for _, v := range inf.valves {
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

	tab := map[state]int{
		{p1: aa, p2: aa}: 0,
	}

	maxrel := 0
	for t := 0; t < timelimit; t++ {
		log.Printf("Commencing timestep %d - considering %d states\n", t, len(tab))

		var srs []staterel
		for s, r := range tab {
			srs = append(srs, staterel{s, r})
		}

		if len(srs) < runtime.NumCPU() {
			out := inf.worker(t, srs)
			tab = merge(out)
			continue
		}

		chunkSize := len(srs) / runtime.NumCPU()
		outch := make(chan []staterel)

		for j := 0; j < runtime.NumCPU(); j++ {
			j := j
			go func() {
				work := srs[j*chunkSize : (j+1)*chunkSize]
				if j == runtime.NumCPU()-1 {
					work = srs[j*chunkSize:]
				}
				outch <- inf.worker(t, work)
			}()
		}

		outs := make([][]staterel, 0, runtime.NumCPU())
		for j := 0; j < runtime.NumCPU(); j++ {
			outs = append(outs, <-outch)
		}
		tab = merge(outs...)
	}

	fmt.Println(maxrel)
}
