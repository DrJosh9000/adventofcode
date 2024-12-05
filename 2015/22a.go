package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2015
// Day 22, part a

// Fiddly!

type state struct {
	hp, mana int

	stimer, ptimer, rtimer int

	bossHP int
}

type spell struct {
	cost  int
	allow func(state) bool
	act   func(state) state
}

func allow(state) bool { return true }

var (
	spells = []*spell{
		{ // magic missile
			cost:  53,
			allow: allow,
			act: func(s state) state {
				s.bossHP -= 4
				return s
			},
		},
		{ // drain
			cost:  73,
			allow: allow,
			act: func(s state) state {
				s.hp += 2
				s.bossHP -= 2
				return s
			},
		},
		{ // shield
			cost: 113,
			allow: func(s state) bool {
				return s.stimer == 0
			},
			act: func(s state) state {
				s.stimer = 6
				return s
			},
		},
		{ // poison
			cost: 173,
			allow: func(s state) bool {
				return s.ptimer == 0
			},
			act: func(s state) state {
				s.ptimer = 6
				return s
			},
		},
		{ // recharge
			cost: 229,
			allow: func(s state) bool {
				return s.rtimer == 0
			},
			act: func(s state) state {
				s.rtimer = 5
				return s
			},
		},
	}
)

func main() {
	var bossHP, bossDmg int
	for _, line := range exp.MustReadLines("inputs/22.txt") {
		s, v, ok := strings.Cut(line, ": ")
		if !ok {
			continue
		}
		n := exp.Must(strconv.Atoi(v))
		switch s {
		case "Hit Points":
			bossHP = n
		case "Damage":
			bossDmg = n
		}
	}

	start := state{
		hp:     50,
		mana:   500,
		bossHP: bossHP,
	}
	mincost := math.MaxInt
	algo.Dijkstra(start, func(s state, d int) (map[state]int, error) {
		// win or loss
		if s.hp <= 0 {
			return nil, nil
		}
		if s.bossHP <= 0 {
			if d < mincost {
				mincost = d
			}
			return nil, nil
		}

		if s.ptimer > 0 {
			s.bossHP -= 3
			s.ptimer--
		}
		if s.stimer > 0 {
			s.stimer--
		}
		if s.rtimer > 0 {
			s.mana += 101
			s.rtimer--
		}

		next := make(map[state]int)

		for _, sp := range spells {
			if sp.cost > s.mana {
				continue
			}
			if !sp.allow(s) {
				continue
			}

			// cast sp & deduct cost
			t := sp.act(s)
			t.mana -= sp.cost

			// boss turn.
			// effects apply both at the start of the player & boss turns...
			if t.ptimer > 0 {
				t.bossHP -= 3
				t.ptimer--
			}
			if t.rtimer > 0 {
				t.mana += 101
				t.rtimer--
			}

			// did we kill the boss at this point?
			if t.bossHP <= 0 {
				next[t] = sp.cost
				continue
			}

			armor := 0
			if t.stimer > 0 {
				armor = 7
				t.stimer--
			}

			bdmg := bossDmg - armor
			if bdmg <= 0 {
				bdmg = 1
			}
			t.hp -= bdmg

			next[t] = sp.cost
		}

		return next, nil
	})

	fmt.Println(mincost)
}
