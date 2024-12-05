package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 21

type item struct{ cost, dmg, arm int }

var (
	// Weapons:    Cost  Damage  Armor
	// Dagger        8     4       0
	// Shortsword   10     5       0
	// Warhammer    25     6       0
	// Longsword    40     7       0
	// Greataxe     74     8       0
	weapons = []*item{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}

	// Armor:      Cost  Damage  Armor
	// Leather      13     0       1
	// Chainmail    31     0       2
	// Splintmail   53     0       3
	// Bandedmail   75     0       4
	// Platemail   102     0       5
	armor = []*item{
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}

	// Rings:      Cost  Damage  Armor
	// Damage +1    25     1       0
	// Damage +2    50     2       0
	// Damage +3   100     3       0
	// Defense +1   20     0       1
	// Defense +2   40     0       2
	// Defense +3   80     0       3
	rings = []*item{
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}
)

func main() {
	var bossHP, bossDmg, bossArm int
	for _, line := range exp.MustReadLines("inputs/21.txt") {
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
		case "Armor":
			bossArm = n
		}
	}

	fight := func(items ...*item) (int, bool) {
		hp, bhp := 100, bossHP
		cost, dmg, arm := 0, 0, 0
		for _, i := range items {
			cost += i.cost
			dmg += i.dmg
			arm += i.arm
		}
		for {
			// player goes first
			d := dmg - bossArm
			if d <= 0 {
				d = 1
			}
			bhp -= d
			if bhp <= 0 {
				return cost, true
			}

			bd := bossDmg - arm
			if bd <= 0 {
				bd = 1
			}
			hp -= bd
			if hp <= 0 {
				return cost, false
			}
		}
	}

	mincost, maxcost := math.MaxInt, 0
	update := func(cost int, win bool) {
		if win && cost < mincost {
			mincost = cost
		}
		if !win && cost > maxcost {
			maxcost = cost
		}
	}

	for _, w := range weapons {
		// all else is optional
		update(fight(w))
		for _, a := range armor {
			// 0 or 1 armor
			update(fight(w, a))
			for _, r1 := range rings {
				// 0-2 rings
				update(fight(w, a, r1))
				for _, r2 := range rings {
					// shop only has one of each item
					if r1 == r2 {
						continue
					}
					update(fight(w, a, r1, r2))
				}
			}
		}
	}

	fmt.Println(mincost, maxcost)
}
