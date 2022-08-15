package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/DrJosh9000/exp"
)

type set[T comparable] map[T]struct{}

func (s set[T]) add(t T)     { s[t] = struct{}{} }
func (s set[T]) ni(t T) bool { _, yes := s[t]; return yes }

type group struct {
	units, hp, atk, ini int
	atktype             string
	weakTo, immuneTo    set[string]
}

func (g *group) effPower() int { return g.units * g.atk }

func (g *group) dmgTo(h *group) int {
	if h.immuneTo.ni(g.atktype) {
		return 0
	}
	if h.weakTo.ni(g.atktype) {
		return g.effPower() * 2
	}
	return g.effPower()
}

func main() {
	var armies [2][]*group
	army := 0
	exp.MustForEachLineIn("inputs/24.txt", func(line string) {
		switch line {
		case "":
			// skip
		case "Immune System:":
			army = 0
		case "Infection:":
			army = 1
		default:
			var g group
			var part3 string
			part1, part2, found := strings.Cut(line, "(")
			if found {
				part2, part3, found = strings.Cut(part2, ")")
				if !found {
					log.Fatalf("Line with '(' but not ')'? %q", line)
				}
				for _, att := range strings.Split(part2, "; ") {
					switch {
					case strings.HasPrefix(att, "weak to"):
						g.weakTo = make(set[string])
						for _, wk := range strings.Split(strings.TrimPrefix(att, "weak to "), ",") {
							wk = strings.TrimSpace(wk)
							g.weakTo.add(wk)
						}
					case strings.HasPrefix(att, "immune to"):
						g.immuneTo = make(set[string])
						for _, imm := range strings.Split(strings.TrimPrefix(att, "immune to "), ",") {
							imm = strings.TrimSpace(imm)
							g.immuneTo.add(imm)
						}
					}
				}
				part1 += strings.TrimSpace(part3)
			}
			if _, err := fmt.Sscanf(part1, "%d units each with %d hit points with an attack that does %d %s damage at initiative %d", &g.units, &g.hp, &g.atk, &g.atktype, &g.ini); err != nil {
				log.Fatalf("Couldn't parse line %q: %v", part1, err)
			}
			armies[army] = append(armies[army], &g)
		}
	})

	for army, a := range armies {
		fmt.Printf("Army %d:\n", army)
		for _, g := range a {
			fmt.Printf("%d units each with %d hp, %d %s atk, %d ini (weak to ", g.units, g.hp, g.atk, g.atktype, g.ini)
			for wk := range g.weakTo {
				fmt.Print(wk)
				fmt.Print(",")
			}
			fmt.Print("; immune to ")
			for imm := range g.immuneTo {
				fmt.Print(imm)
				fmt.Print(",")
			}
			fmt.Println(")")
		}
	}

	for { // fights
		// Target selection
		targ := make(map[*group]*group)
		for army, a := range armies {
			sort.Slice(a, func(i, j int) bool {
				// Higher effective power goes first
				ip, jp := a[i].effPower(), a[j].effPower()
				if ip == jp {
					// ...in a tie, higher initiative.
					return a[i].ini > a[j].ini
				}
				return ip > jp
			})
			// Track selected targets to ensure no group is targetted twice
			chosen := make(set[*group])
			for _, g := range a {
				if g.units <= 0 {
					break
				}
				var best *group
				bestDmg, bestPow, bestIni := 0, 0, 0
				for _, h := range armies[1-army] {
					if h.units <= 0 || chosen.ni(h) {
						continue
					}
					dmg := g.dmgTo(h)
					if dmg == 0 {
						continue
					}
					if dmg > bestDmg {
						best = h
						bestDmg = dmg
						bestPow = h.effPower()
						bestIni = h.ini
						continue
					}
					if dmg != bestDmg {
						continue
					}
					// dmg == bestDmg
					ep := h.effPower()
					if ep > bestPow {
						best = h
						bestPow = ep
						bestIni = h.ini
						continue
					}
					if ep != bestPow {
						continue
					}
					// ep == bestPow
					if h.ini > bestIni {
						best = h
						bestIni = h.ini
					}
				}
				if best != nil {
					targ[g] = best
					chosen.add(best)
				}
			}
		}

		// Attacking
		gs := make([]*group, 0, len(targ))
		for g := range targ {
			gs = append(gs, g)
		}
		sort.Slice(gs, func(i, j int) bool {
			return gs[i].ini > gs[j].ini
		})

		for _, g := range gs {
			if g.units <= 0 {
				continue
			}
			h := targ[g]
			if h.units <= 0 {
				continue
			}
			h.units -= g.dmgTo(h) / h.hp
		}

		var units [2]int
		for i, a := range armies {
			for _, g := range a {
				if g.units < 0 {
					continue
				}
				units[i] += g.units
			}
		}

		if units[0] == 0 || units[1] == 0 {
			fmt.Println("Army 0 has", units[0], "units")
			fmt.Println("Army 1 has", units[1], "units")
			break
		}
	}
}
