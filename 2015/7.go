package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 7

var wires map[string]uint16

func eval(s string) uint16 {
	n, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return wires[s]
	}
	return uint16(n)
}

type unary struct {
	in string
}

func (g unary) id() uint16  { return eval(g.in) }
func (g unary) not() uint16 { return ^eval(g.in) }

type binary struct {
	a, b string
}

func (g binary) and() uint16 { return eval(g.a) & eval(g.b) }
func (g binary) or() uint16  { return eval(g.a) | eval(g.b) }
func (g binary) lsh() uint16 { return eval(g.a) << eval(g.b) }
func (g binary) rsh() uint16 { return eval(g.a) >> eval(g.b) }

type gate struct {
	out string
	op  func() uint16
}

func main() {
	var circuit []gate
	var b int
	for _, line := range exp.MustReadLines("inputs/7.txt") {
		tokens := strings.Fields(line)
		var g gate
		switch {
		case tokens[1] == "->":
			g = gate{
				out: tokens[2],
				op:  unary{tokens[0]}.id,
			}
		case tokens[0] == "NOT":
			g = gate{
				out: tokens[3],
				op:  unary{tokens[1]}.not,
			}
		case tokens[1] == "AND":
			g = gate{
				out: tokens[4],
				op:  binary{tokens[0], tokens[2]}.and,
			}
		case tokens[1] == "OR":
			g = gate{
				out: tokens[4],
				op:  binary{tokens[0], tokens[2]}.or,
			}
		case tokens[1] == "LSHIFT":
			g = gate{
				out: tokens[4],
				op:  binary{tokens[0], tokens[2]}.lsh,
			}
		case tokens[1] == "RSHIFT":
			g = gate{
				out: tokens[4],
				op:  binary{tokens[0], tokens[2]}.rsh,
			}
		default:
			log.Fatalf("Unknown form %q", line)
		}
		if g.out == "b" {
			b = len(circuit)
		}
		circuit = append(circuit, g)
	}

	// This is day 7. Don't overthink it.
	sim := func() {
		wires = make(map[string]uint16)
		change := true
		for change {
			change = false
			for _, g := range circuit {
				if o, n := wires[g.out], g.op(); o != n {
					wires[g.out] = n
					change = true
				}
			}
		}
	}

	sim()
	a := wires["a"]
	fmt.Println(a)

	circuit[b] = gate{out: "b", op: func() uint16 { return a }}
	sim()
	fmt.Println(wires["a"])
}
