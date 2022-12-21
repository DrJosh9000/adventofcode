package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"

	"github.com/mitchellh/go-z3"
)

// Advent of Code 2022
// Day 21, part b

func main() {
	cfg := z3.NewConfig()
	ctx := z3.NewContext(cfg)
	cfg.Close()
	defer ctx.Close()

	s := ctx.NewSolver()
	defer s.Close()

	vars := make(map[string]*z3.AST)
	input := exp.MustReadLines("inputs/21.txt")
	for _, line := range input {
		name, _, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalf("Couldn't find cut string in %q", line)
		}

		vars[name] = ctx.Const(ctx.Symbol(name), ctx.IntSort())
	}

	for _, line := range input {
		name, op, _ := strings.Cut(line, ": ")
		if name == "humn" {
			continue
		}

		opf := strings.Fields(op)

		x := vars[name]
		if len(opf) == 1 {
			s.Assert(x.Eq(ctx.Int(exp.Must(strconv.Atoi(opf[0])), ctx.IntSort())))
			continue
		}

		l, o, r := vars[opf[0]], opf[1], vars[opf[2]]
		if name == "root" {
			s.Assert(l.Eq(r))
			continue
		}

		switch o {
		case "+":
			s.Assert(x.Eq(l.Add(r)))
		case "-":
			s.Assert(x.Eq(l.Sub(r)))
		case "*":
			s.Assert(x.Eq(l.Mul(r)))
		case "/":
			s.Assert(l.Eq(x.Mul(r)))
		}
	}

	if v := s.Check(); v != z3.True {
		log.Fatal("Unsolvable")
	}

	m := s.Model()
	a := m.Assignments()
	m.Close()
	fmt.Println(a["humn"])
}
