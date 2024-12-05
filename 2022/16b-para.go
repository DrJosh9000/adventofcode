package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"drjosh.dev/exp"
)

// Advent of Code 2022
// Day 16, part b

type valve struct {
	name    string
	rate    int32
	tunsrc  []string
	tunnels []int
}

type state uint

// 6 + 6 + 15 = 27

func (s state) unpack() (p1, p2 int, open uint) {
	return int(s) & 0x3f, (int(s) >> 6) & 0x3f, uint(s) >> 12
}

func pack(p1, p2 int, open uint) state {
	return state(open)<<12 | state(p2)<<6 | state(p1)
}

func main() {
	cpuf := exp.Must(os.Create("cpu-16b-para.prof"))
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

	tab := make([]int32, 1<<27)
	for i := range tab {
		tab[i] = -1
	}
	tab[pack(aa, aa, 0)] = 0
	t2 := make([]int32, 1<<27)

	var maxrel int32
	upmax := func(r int32) {
		for {
			or := atomic.LoadInt32(&maxrel)
			if or >= r {
				break
			}
			if atomic.CompareAndSwapInt32(&maxrel, or, r) {
				break
			}
		}
	}

	for t := int32(0); t < timelimit; t++ {
		log.Printf("Commencing timestep %d - considering %d states\n", t, len(tab))
		var valid int64

		for i := range t2 {
			t2[i] = -1
		}

		upd := func(s state, r int32) {
			for {
				or := atomic.LoadInt32(&t2[s])
				if or >= r {
					break
				}
				if atomic.CompareAndSwapInt32(&t2[s], or, r) {
					break
				}
			}
		}

		N := runtime.NumCPU()
		cs := len(tab) / N
		var wg sync.WaitGroup
		for i := 0; i < N; i++ {
			// Splitting the work into contiguous chunks has better performance
			// than splitting the states by modulus N.
			start := i * cs
			chunk := tab[start:]
			if i < N-1 {
				chunk = chunk[:cs]
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				var val int64
				var mrel int32
				for si, r := range chunk {
					if r < 0 {
						continue
					}
					val++
					s := state(si + start)
					p1, p2, open := s.unpack()
					v1 := valves[p1]
					v2 := valves[p2]

					// Both move
					for _, j := range v1.tunnels {
						for _, k := range v2.tunnels {
							upd(pack(j, k, open), r)
						}
					}

					if v1.rate > 0 && open&(1<<p1) == 0 {
						// p1 has the option of opening a valve, while p2 moves
						nr := r + (timelimit-t-1)*v1.rate
						if nr > mrel {
							mrel = nr
						}
						if open := open | 1<<p1; open != allopen {
							for _, j := range v2.tunnels {
								upd(pack(p1, j, open), nr)
							}
						}
					}

					if v2.rate > 0 && open&(1<<p2) == 0 {
						// p2 has the option of opening a valve, while p1 moves.
						nr := r + (timelimit-t-1)*v2.rate
						if nr > mrel {
							mrel = nr
						}

						if open := open | 1<<p2; open != allopen {
							for _, j := range v1.tunnels {
								upd(pack(j, p2, open), nr)
							}
						}
					}

					// also consider both opening, if that's possible
					// ... make sure they are at separate valves!
					if v1.rate > 0 && open&(1<<p1) == 0 && v2.rate > 0 && open&(1<<p2) == 0 && p1 != p2 {
						nr := r + (timelimit-t-1)*(v1.rate+v2.rate)
						if nr > mrel {
							mrel = nr
						}

						if open := open | 1<<p1 | 1<<p2; open != allopen {
							upd(pack(p1, p2, open), nr)
						}
					}
				}

				atomic.AddInt64(&valid, val)
				upmax(mrel)
			}()
		}

		wg.Wait()

		log.Printf("Ending timestep %d - considered %d valid states", t, valid)

		tab, t2 = t2, tab
	}

	fmt.Println(maxrel)
}
