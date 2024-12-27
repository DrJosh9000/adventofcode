package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2024
// Day 24, part b

const inputPath = "2024/inputs/24.txt"

func main() {
	type gate struct {
		x, y, z, op string
	}

	outputRE := regexp.MustCompile(`z\d\d`)

	lines := exp.MustReadLines(inputPath)
	wires := make(algo.Set[string])
	// var gates []gate
	byOut := make(map[string]*gate)
	byIn := make(map[string]algo.Set[*gate])
	for _, line := range lines {
		var g gate
		if exp.Smatchf(line, "%s %s %s -> %s", &g.x, &g.op, &g.y, &g.z) {
			wires.Insert(g.x, g.y, g.z)
			byOut[g.z] = &g
			byIn[g.x] = byIn[g.x].Insert(&g)
			byIn[g.y] = byIn[g.y].Insert(&g)
		}
	}

	sus := make(algo.Set[string])

	role := make(map[string]string)

	// Each pair of input lines should connect to a half-adder (one XOR and one
	// AND)
	for i := range 45 {
		xi, yi := fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i)
		xg := byIn[xi]
		if !xg.Equal(byIn[yi]) {
			// panic("different gates connected to inputs (not half-adder)! " + xi + " + " + yi)
			sus.Insert(xi, yi)
		}
		if len(xg) != 2 {
			// panic("wrong number of gates connected to inputs! " + xi + " + " + yi)
			sus.Insert(xi, yi)
		}
		want := algo.MakeSet("AND", "XOR")
		for g := range xg {
			switch g.op {
			case "XOR":
				// q from the first half-adder should not
				// be connected to any output (except z00 = x00 ^ y00)
				if outputRE.MatchString(g.z) && g.z != "z00" {
					sus.Insert(g.z)
				}
				r := fmt.Sprintf("ha1q%02d", i)
				if role[r] != "" && role[r] != g.z {
					log.Printf("role %s already assigned to wire %s while trying to assign to %s", r, role[r], g.z)
					sus.Insert(g.z)
				}
				role[r] = g.z
			case "AND":
				// c from the first half-adder should not
				// be connected to any output
				if outputRE.MatchString(g.z) {
					sus.Insert(g.z)
				}
				r := fmt.Sprintf("ha1c%02d", i)
				if role[r] != "" && role[r] != g.z {
					log.Printf("role %s already assigned to wire %s while trying to assign to %s", r, role[r], g.z)
					sus.Insert(g.z)
				}
				role[r] = g.z
			case "OR":
				log.Printf("OR connected within in half-adder %s, %s", xi, yi)
				sus.Insert(xi, yi)
			}
			delete(want, g.op)
		}
		if len(want) != 0 {
			log.Printf("two gates of same op connected to inputs %s, %s", xi, yi)
			sus.Insert(xi, yi)
		}
	}

	// Each output line should be driven by an XOR of a second-stage half-adder,
	// (except the last one, which is driven by the final carry).
	for i := range 45 { // z45 appears to be correct on the graph
		zi := fmt.Sprintf("z%02d", i)
		g := byOut[zi]
		if g.op != "XOR" {
			sus.Insert(zi)
		}
		r := fmt.Sprintf("ha2q%02d", i)
		if role[r] != "" && role[r] != g.z {
			log.Printf("role %s already assigned to wire %s", r, g.z)
			sus.Insert(g.z)
		}
		role[r] = g.z
		// If i > 0 then the inputs to this XOR should be driving a second-stage
		// half-adder.
		if i == 0 {
			continue
		}
		xg := byIn[g.x]
		if !xg.Equal(byIn[g.y]) {
			log.Printf("different gate sets for %s, %s", g.x, g.y)
			sus.Insert(g.x, g.y)
		}
		if len(xg) != 2 {
			log.Printf("wrong number of gates for second-stage half-adder %s, %s", g.x, g.y)
			sus.Insert(g.x, g.y)
		}

		want := algo.MakeSet("AND", "XOR")
		for ga := range xg {
			switch g.op {
			case "XOR":
				// already handled
			case "AND":
				// c from the second half-adder should be connected to an OR,
				// not any output
				if outputRE.MatchString(ga.z) {
					log.Printf("carry of second-stage half adder connected to %s", ga.z)
					sus.Insert(ga.z)
				}
				r := fmt.Sprintf("ha2c%02d", i)
				if role[r] != "" && role[r] != ga.z {
					log.Printf("role %s already assigned to wire %s while trying to assign to %s", r, role[r], ga.z)
					sus.Insert(ga.z)
				}
				role[r] = ga.z

				cgs := byIn[ga.z]
				if len(cgs) != 1 {
					log.Printf("second-stage half-adder carry %d connected to zero or multiple gates", i)
					sus.Insert(ga.z)
					break
				}
				cg := cgs.Any()
				if cg.op != "OR" {
					log.Printf("second-stage half-adder carry %d not connected to OR", i)
					sus.Insert(ga.z)
				}

				// The other input to the OR should be a first-stage carry!

			case "OR":
				log.Printf("OR connected within in half-adder %s, %s", ga.x, ga.y)
				sus.Insert(ga.x, ga.y)
			}
			delete(want, ga.op)
		}
		if len(want) != 0 {
			log.Printf("two gates of same op connected to inputs %s, %s", g.x, g.y)
			sus.Insert(g.x, g.y)
		}
	}

	func() {
		// The second half-adder between bits 0 and 1 is connected without an OR
		// because there is no second half-adder before bit 0.
		xi, yi := role["ha1q01"], role["ha1c00"]
		xg := byIn[xi]
		if !xg.Equal(byIn[yi]) {
			log.Printf("different gate sets for ha1q01=%s, ha1c00=%s", xi, yi)
		}
		if len(xg) != 2 {
			log.Printf("wrong number of gates for ha1q01=%s, ha1c00=%s", xi, yi)
		}
		want := algo.MakeSet("AND", "XOR")
		for g := range xg {
			switch g.op {
			case "XOR":
				// q from the second half adder should be the z01 output
				if g.z != "z01" {
					log.Print("z01 is not the output of the second half-adder")
					sus.Insert(g.z)
				}
				r := "ha2q01"
				if role[r] != "" && role[r] != g.z {
					log.Printf("role %s already assigned to wire %s while trying to assign to %s", r, role[r], g.z)
					sus.Insert(g.z)
				}
				role[r] = g.z
			case "AND":
				// c from the second-stage half adder should only
				// be connected to the carry-OR, not an output, including
				// the last output bit.
				if outputRE.MatchString(g.z) {
					sus.Insert(g.z)
				}
				if gz := byIn[g.z]; len(gz) != 1 || gz.Any().op != "OR" {
					sus.Insert(g.z)
				}
				r := "ha2c01"
				if role[r] != "" && role[r] != g.z {
					log.Printf("role %s already assigned to wire %s while trying to assign to %s", r, role[r], g.z)
					sus.Insert(g.z)
				}
				role[r] = g.z

				cgs := byIn[g.z]
				if len(cgs) != 1 {
					log.Printf("second-stage half-adder carry 1 connected to zero or multiple gates")
					sus.Insert(g.z)
					break
				}
				cg := cgs.Any()
				if cg.op != "OR" {
					log.Printf("second-stage half-adder carry 1 not connected to OR")
					sus.Insert(g.z)
				}

			case "OR":
				log.Printf("OR connected within in half-adder %s, %s", xi, yi)
				sus.Insert(xi, yi)
			}
			delete(want, g.op)
		}
		if len(want) != 0 {
			log.Printf("two gates of same op connected to inputs %s, %s", xi, yi)
			sus.Insert(xi, yi)
		}
	}()

	// // The two carries should be OR-ed together
	// for i := 1; ; i++ {
	// 	c1, c2 := fmt.Sprintf("ha1c%02d", i), fmt.Sprintf("ha2c%02d", i)
	// 	xi, yi := role[c1], role[c2]
	// 	if _, has := wires[xi]; !has {
	// 		break
	// 	}
	// 	xg := byIn[xi]
	// 	if !xg.Equal(byIn[yi]) {
	// 		//panic(fmt.Sprintf("different gates connected! %s=%s + %s=%s", c1, xi, c2, yi))
	// 		continue
	// 	}
	// 	if len(xg) != 1 {
	// 		//panic("wrong number of gates connected! " + xi + " + " + yi)
	// 		continue
	// 	}
	// 	g := xg.Any()
	// 	if g.op != "OR" {
	// 		sus.Insert(xi, yi)
	// 	}
	// 	// c := fmt.Sprintf("carry%02d", i)
	// }

	// fmt.Println(role)
	suss := sus.ToSlice()
	slices.Sort(suss)
	fmt.Println(strings.Join(suss, ","))

	// fmt.Println("digraph {")
	// for w := range wires {
	// 	fmt.Printf("%s [shape=plain];\n", w)
	// }
	// for _, g := range gates {
	// 	fmt.Printf("gate_%s [shape=box,label=%s];\n", g.z, g.op)
	// 	fmt.Printf("%s -> gate_%s;\n", g.x, g.z)
	// 	fmt.Printf("%s -> gate_%s;\n", g.y, g.z)
	// 	fmt.Printf("gate_%s -> %s;\n", g.z, g.z)
	// }
	// fmt.Println("}")
}

// frn,gmq,vtj,wnf,wtt,z05,z21,z39
