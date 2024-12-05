package main

import (
	"fmt"
	"strings"
	"sync/atomic"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/para"
)

// Advent of Code 2023
// Day 19, part b

const inputPath = "2023/inputs/19.txt"

type rule struct {
	varn, cmp, dest string
	thres           int
}

type part struct {
	wf         string
	x, m, a, s []algo.Range[int]
}

func rlen(r algo.Range[int]) int { return r.Max - r.Min + 1 }

func (r *rule) act(p part) (acc, pass part) {
	acc, pass = p, p

	var v, vacc, vpass *[]algo.Range[int]
	switch r.varn {
	case "":
		return p, part{}
	case "x":
		v, vacc, vpass = &p.x, &acc.x, &pass.x
	case "m":
		v, vacc, vpass = &p.m, &acc.m, &pass.m
	case "a":
		v, vacc, vpass = &p.a, &acc.a, &pass.a
	case "s":
		v, vacc, vpass = &p.s, &acc.s, &pass.s
	default:
		panic("unknown variable " + r.varn)
	}

	var sub algo.Range[int]
	switch r.cmp {
	case "<":
		sub = algo.NewRange(r.thres, 4000)

	case ">":
		sub = algo.NewRange(1, r.thres)

	default:
		panic("unknown comparison " + r.cmp)
	}

	var va, vp []algo.Range[int]
	for _, vr := range *v {
		va = append(va, algo.RangeSubtract(vr, sub)...)
		in := vr.Intersection(sub)
		if !in.IsEmpty() {
			vp = append(vp, in)
		}
	}
	*vacc = va
	*vpass = vp
	return acc, pass
}

func main() {
	lines := exp.MustReadLines(inputPath)

	byName := make(map[string][]rule)
	for _, line := range lines {
		if line == "" {
			break
		}
		var wf []rule
		name, body, ok := strings.Cut(line, "{")
		if !ok {
			panic("invalid workflow " + line)
		}
		for _, rs := range strings.Split(strings.TrimSuffix(body, "}"), ",") {
			cond, dest, ok := strings.Cut(rs, ":")
			var r rule
			if ok {
				r = rule{
					varn:  cond[:1],
					cmp:   cond[1:2],
					dest:  dest,
					thres: exp.MustAtoi(cond[2:]),
				}
			} else {
				r = rule{dest: cond}
			}
			wf = append(wf, r)
		}
		byName[name] = wf
	}

	full := []algo.Range[int]{algo.NewRange(1, 4000)}
	var sum atomic.Int64
	q := para.NewQueue(part{
		wf: "in",
		x:  full,
		m:  full,
		a:  full,
		s:  full,
	})
	q.Process(func(p part) {
		wf, ok := byName[p.wf]
		if !ok {
			panic("unknown workflow " + p.wf)
		}
		for _, r := range wf {
			if len(p.x) == 0 || len(p.m) == 0 || len(p.a) == 0 || len(p.s) == 0 {
				return
			}

			acc, pass := r.act(p)

			switch r.dest {
			case "A":
				x := algo.Sum(algo.Map(acc.x, rlen))
				m := algo.Sum(algo.Map(acc.m, rlen))
				a := algo.Sum(algo.Map(acc.a, rlen))
				s := algo.Sum(algo.Map(acc.s, rlen))
				sum.Add(int64(x * m * a * s))

			case "R":
				// nop

			default:
				acc.wf = r.dest
				q.Push(acc)
			}

			p = pass
		}
	})

	fmt.Println(sum.Load())
}
