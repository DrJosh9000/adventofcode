package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"sort"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

var teams = map[byte]string{
	'E': "Elves",
	'G': "Goblins",
}

var adjacent = []image.Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

func isAdjacent(p, q image.Point) bool {
	d := p.Sub(q)
	return (d.X == 0 && (d.Y == -1 || d.Y == 1)) ||
		(d.Y == 0 && (d.X == -1 || d.X == 1))
}

type unit struct {
	pt image.Point
	team byte
	hp, atk int
}

func (u *unit) alive() bool {
	return u.hp > 0
}

func (u *unit) String() string {
	return fmt.Sprintf("%c%v(hp=%d)", u.team, u.pt, u.hp)
}

// Unit takes a turn. Returns true if combat should continue.
func (u *unit) takeTurn(grid [][]byte, units []*unit) bool {
	if !u.alive() {
		// It's dead.
		return true
	}
	
	// Determine targets.
	var targets, attack []*unit
	for _, v := range units {
		if !v.alive() || u.team == v.team {
			// Don't target/attack dead units or units on the same team.
			continue
		}
		targets = append(targets, v)
		if isAdjacent(u.pt, v.pt) {
			attack = append(attack, v)
		}
	}
	
	if len(targets) == 0 {
		// No targets remain - combat ends.
		return false
	}
	
	// If none are in range already, do a movement first.
	if len(attack) == 0 {
		dest, err := bestAdjacent(grid, u.pt, targets)
		if err != nil {
			// No reachable targets; turn ends.
			return true
		}
		
		// Find best step to take. This is just a search through the
		// reverse paths (from dest back to squares next to u).
		next, err := bestAdjacent(grid, dest, []*unit{u})
		if err != nil {
			// That's unexpected!
			log.Fatalf("Couldn't find a reverse path when a forward path was already found?? %v", err)
		}
		
		// Take the step.
		grid[u.pt.Y][u.pt.X] = '.'
		grid[next.Y][next.X] = u.team
		u.pt = next
		
		// Okay, are any targets in range now?
		for _, t := range targets {
			if isAdjacent(u.pt, t.pt) {
				attack = append(attack, t)
			}
		}
	}
	
	if len(attack) == 0 {
		// Moved but still not in range. End of turn.
		return true
	}
	
	// Attack.
	// Pick target with lowest hp; among equals, first in reading order.
	targ := attack[0]
	for _, t := range attack[1:] {
		if t.hp < targ.hp || (t.hp == targ.hp && (t.pt.Y < targ.pt.Y || t.pt.Y == targ.pt.Y && t.pt.X < targ.pt.X)) {
			targ = t
		}
	}
	
	// Attack!
	targ.hp -= u.atk
	if !targ.alive() {
		// It died!
		grid[targ.pt.Y][targ.pt.X] = '.'
	}
	
	return true
}


// Determine the order of turns based on reading order. Remove any dead units.
func sortUnits(s []*unit) []*unit {
	sort.Slice(s, func(i, j int) bool {
		if s[i].alive() == s[j].alive() {
			if s[i].pt.Y == s[j].pt.Y {
				return s[i].pt.X < s[j].pt.X
			}
			return s[i].pt.Y < s[j].pt.Y
		}
		return s[i].alive() // => !s[j].alive()
	})
	
	// Cull dead units.
	for i, u := range s {
		if u.alive() {
			continue
		}
		s = s[:i]
		break
	}
	return s
}

// Find the "best" adjacent square to any of the given targets. Best means
// shortest path from src; among equal lengths, the first in reading order.
func bestAdjacent(grid [][]byte, src image.Point, targets []*unit) (image.Point, error) {
	inrange := make(map[image.Point]int)
	for _, t := range targets {
		for _, d := range adjacent {
			p := t.pt.Add(d)
			if grid[p.Y][p.X] == '.' {
				inrange[p] = math.MaxInt
			}
		}
	}
	
	// Find distances to in-range points
	algo.FloodFill(src, func(p image.Point, dist int) ([]image.Point, error) {
		var out []image.Point
		if dist < inrange[p] {
			inrange[p] = dist
		}
		for _, d := range adjacent {
			q := p.Add(d)
			if grid[q.Y][q.X] == '.' {
				out = append(out, q)
			}
		}
		return out, nil
	})
	
	// Choose the destination (nearest in-range point; of those,
	// the first in reading order).
	var dest image.Point
	dist := math.MaxInt
	for p, d := range inrange {
		if d < dist || (d == dist && (p.Y < dest.Y || p.Y == dest.Y && p.X < dest.X)) {
			dest, dist = p, d
		}
	}
	
	if dist == math.MaxInt {
		// No targets reachable - end turn.
		return image.Point{}, fmt.Errorf("no path")
	}
	
	return dest, nil
}

func main() {
	var grid [][]byte
	var units []*unit
	j := 0
	exp.MustForEachLineIn("inputs/15.txt", func(line string) {
		row := []byte(line)
		for i, c := range row {
			if c == 'E' || c == 'G' {
				units = append(units, &unit{
					pt: image.Pt(i, j),
					team: c,
					hp: 200,
					atk: 3,
				})
			}
		}
		grid = append(grid, row)
		j++
	})
	
	for round := 0; ; round++ {
		// Determine turn order and cull dead units.
		units = sortUnits(units)
		
		// Each unit takes a turn.
		for _, u := range units {
			if u.takeTurn(grid, units) {
				continue
			}
			// Combat ends.
			for _, r := range grid {
				fmt.Println(string(r))
			}
			hpsum := 0
			for _, u := range units {
				fmt.Println(u)
				if !u.alive() {
					continue
				}
				hpsum += u.hp
			}
			fmt.Println("Combat ends after", round, "full rounds")
			fmt.Println(teams[u.team], "win with", hpsum, "total hit points left")
			fmt.Println("Outcome:", round, "*", hpsum, "=", round * hpsum)
			return
		}
	}
}