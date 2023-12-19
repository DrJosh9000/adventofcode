package main

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/para"
)

// Advent of Code 2023
// Day 19, part a

const inputPath = "2023/inputs/19.txt"

type rule struct {
	varn, cmp, dest string
	thres           int
}

type part struct {
	wf         string
	x, m, a, s int
}

func (r *rule) test(p part) bool {
	var v int
	switch r.varn {
	case "":
		return true
	case "x":
		v = p.x
	case "m":
		v = p.m
	case "a":
		v = p.a
	case "s":
		v = p.s
	default:
		panic("unknown variable " + r.varn)
	}
	switch r.cmp {
	case "<":
		return v < r.thres
	case ">":
		return v > r.thres
	default:
		panic("unknown comparison " + r.cmp)
	}
}

func main() {
	lines := exp.MustReadLines(inputPath)

	var parts []part
	byName := make(map[string][]rule)
	parsingWFs := true
	for _, line := range lines {
		if line == "" && parsingWFs {
			parsingWFs = false
			continue
		}
		if parsingWFs {
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
			continue
		}
		// parsing parts
		var p part
		exp.MustSscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &p.x, &p.m, &p.a, &p.s)
		p.wf = "in"
		parts = append(parts, p)
	}

	var sum atomic.Int64
	q := para.NewQueue(parts...)
	q.Process(func(p part) {
		wf, ok := byName[p.wf]
		if !ok {
			panic("unknown workflow " + p.wf)
		}
		for _, r := range wf {
			if !r.test(p) {
				continue
			}
			switch r.dest {
			case "A":
				sum.Add(int64(p.x + p.m + p.a + p.s))
				return
			case "R":
				// drop
				return
			default:
				p.wf = r.dest
				q.Push(p)
				return
			}
		}
		panic(fmt.Sprintf("part %v didn't match any rule in workflow %v", p, wf))
	})

	fmt.Println(sum.Load())
}
