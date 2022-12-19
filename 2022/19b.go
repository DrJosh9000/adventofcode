package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 19, part b

type blueprint struct {
	num         int
	oreOreCost  int
	clayOreCost int
	obsOreCost  int
	obsClayCost int
	geoOreCost  int
	geoObsCost  int
}

type state struct {
	ore, clay, obs, geo                 int
	orebots, claybots, obsbots, geobots int
}

func (s state) collect() state {
	s.ore += s.orebots
	s.clay += s.claybots
	s.obs += s.obsbots
	s.geo += s.geobots
	return s
}

func main() {
	var bps []blueprint
	for _, line := range exp.MustReadLines("inputs/19.txt") {
		var bp blueprint
		exp.Must(fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.num, &bp.oreOreCost, &bp.clayOreCost, &bp.obsOreCost, &bp.obsClayCost, &bp.geoOreCost, &bp.geoObsCost))

		bps = append(bps, bp)
		if len(bps) == 3 {
			break
		}
	}

	var wg sync.WaitGroup
	geodes := []int{1, 1, 1}
	for i, bp := range bps {
		i, bp := i, bp
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("Evaluating blueprint %d...", bp.num)
			gbt := make([]int, 32)
			algo.FloodFill(state{orebots: 1}, func(s state, t int) ([]state, error) {
				if t == 32 {
					if s.geo > geodes[i] {
						geodes[i] = s.geo
					}
					return nil, nil
				}
				// yeah ok
				if s.geo < gbt[t] {
					return nil, nil
				}
				if s.geo > gbt[t] {
					gbt[t] = s.geo
				}
				next := []state{s.collect()}
				if s.ore >= bp.oreOreCost {
					ns := s.collect()
					ns.ore -= bp.oreOreCost
					ns.orebots++
					next = append(next, ns)
				}
				if s.ore >= bp.clayOreCost {
					ns := s.collect()
					ns.ore -= bp.clayOreCost
					ns.claybots++
					next = append(next, ns)
				}
				if s.ore >= bp.obsOreCost && s.clay >= bp.obsClayCost {
					ns := s.collect()
					ns.ore -= bp.obsOreCost
					ns.clay -= bp.obsClayCost
					ns.obsbots++
					next = append(next, ns)
				}
				if s.ore >= bp.geoOreCost && s.obs >= bp.geoObsCost {
					ns := s.collect()
					ns.ore -= bp.geoOreCost
					ns.obs -= bp.geoObsCost
					ns.geobots++
					next = append(next, ns)
				}
				return next, nil
			})
			log.Printf("Blueprint %d: obtained maximum geodes %d", bp.num, geodes[i])
		}()
	}
	wg.Wait()

	fmt.Println(algo.Prod(geodes))
}
