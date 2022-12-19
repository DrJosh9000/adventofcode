package main

import (
	"fmt"
	"log"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 19, part b

type blueprint struct {
	num         int8
	oreOreCost  int8
	clayOreCost int8
	obsOreCost  int8
	obsClayCost int8
	geoOreCost  int8
	geoObsCost  int8
}

// Each bot costs either 2 or 4 ore.
// If we have 4 orebots, then we have satisfied the ore requirement for any
// bot... otherwise, track ore quantity between 0 and 4 (3 bits)
// Then there are obisidian bots and geode bots...
// Obisidan bots need up to 18 clay, so that's 5 bits
// Geode bots need up to 13 obsidian, so that's 4 bits

type state struct {
	ore, clay, obs                      int8 // 3+5+4 = 12
	orebots, claybots, obsbots, geobots int8 // 4*5 = 20
}

func unpack(s uint32) state {
	return state{
		ore:      int8(s >> 29),
		clay:     int8((s >> 24) & 0b11111),
		obs:      int8((s >> 20) & 0b1111),
		orebots:  int8((s >> 15) & 0b11111),
		claybots: int8((s >> 10) & 0b11111),
		obsbots:  int8((s >> 5) & 0b11111),
		geobots:  int8(s & 0b11111),
	}
}

func (s state) pack() uint32 {
	return (uint32(s.ore) << 29) |
		(uint32(s.clay&0b11111) << 24) |
		(uint32(s.obs&0b1111) << 20) |
		(uint32(s.claybots&0b11111) << 10) |
		(uint32(s.obsbots&0b11111) << 5) |
		uint32(s.geobots&0b11111)
}

func (s state) collect() state {
	s.ore += s.orebots
	if s.ore > 7 {
		s.ore = 7
	}
	s.clay += s.claybots
	if s.clay > 31 {
		s.clay = 31
	}
	s.obs += s.obsbots
	if s.obs > 15 {
		s.obs = 15
	}
	return s
}

func main() {
	var bps []blueprint
	for _, line := range exp.MustReadLines("inputs/19test.txt") {
		var bp blueprint
		exp.Must(fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.num, &bp.oreOreCost, &bp.clayOreCost, &bp.obsOreCost, &bp.obsClayCost, &bp.geoOreCost, &bp.geoObsCost))

		bps = append(bps, bp)
		if len(bps) == 3 {
			break
		}
	}

	// Eight gigabytes!~
	t1 := make([]int8, 1<<32)
	t2 := make([]int8, 1<<32)

	geodes := []int8{1, 1, 1}
	for i, bp := range bps {
		log.Printf("Evaluating blueprint %d...", bp.num)

		for i := range t1 {
			t1[i] = -1
		}
		t1[state{orebots: 1}.pack()] = 0

		for t := 0; t < 32; t++ {
			log.Printf("Beginning timestep %d...", t)
			for si, geo := range t1 {
				if geo < 0 {
					continue
				}
				s := unpack(uint32(si))

				ns := s.collect().pack()
				if g2 := geo + s.geobots; g2 > t2[ns] {
					t2[ns] = g2
				}

				if s.ore >= bp.oreOreCost || s.orebots >= bp.oreOreCost {
					ns := s.collect()
					ns.ore -= bp.oreOreCost
					ns.orebots++
					ps := ns.pack()
					if g2 := geo + s.geobots; g2 > t2[ps] {
						t2[ps] = g2
					}
				}
				if s.ore >= bp.clayOreCost || s.orebots >= bp.clayOreCost {
					ns := s.collect()
					ns.ore -= bp.clayOreCost
					ns.claybots++
					ps := ns.pack()
					if g2 := geo + s.geobots; g2 > t2[ps] {
						t2[ps] = g2
					}
				}
				if (s.ore >= bp.obsOreCost || s.orebots >= bp.obsOreCost) && (s.clay >= bp.obsClayCost || s.claybots >= bp.obsClayCost) {
					ns := s.collect()
					ns.ore -= bp.obsOreCost
					ns.clay -= bp.obsClayCost
					ns.obsbots++
					ps := ns.pack()
					if g2 := geo + s.geobots; g2 > t2[ps] {
						t2[ps] = g2
					}
				}
				if (s.ore >= bp.geoOreCost || s.orebots >= bp.geoOreCost) && (s.obs >= bp.geoObsCost || s.orebots >= bp.geoObsCost) {
					ns := s.collect()
					ns.ore -= bp.geoOreCost
					ns.obs -= bp.geoObsCost
					ns.geobots++
					ps := ns.pack()
					if g2 := geo + s.geobots; g2 > t2[ps] {
						t2[ps] = g2
					}
				}
			}
			t1, t2 = t2, t1
		}
		for _, geo := range t1 {
			if geo > geodes[i] {
				geodes[i] = geo
			}
		}

		log.Printf("Blueprint %d: obtained maximum geodes %d", bp.num, geodes[i])
	}
	fmt.Println(algo.Prod(geodes))
}
