package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 7, part b

type entry interface {
	size() int
}

type file struct {
	n string
	s int
}

func (f *file) size() int { return f.s }

type dir struct {
	n string
	c map[string]entry
	p *dir
	s int
}

func (d *dir) size() int {
	if d.s > 0 {
		return d.s
	}
	s := 0
	for _, e := range d.c {
		s += e.size()
	}
	d.s = s
	return s
}

func main() {
	root := &dir{
		c: make(map[string]entry),
	}
	root.p = root
	pwd := root
	cmd := ""
	for _, line := range exp.MustReadLines("inputs/7.txt") {
		switch {
		case strings.HasPrefix(line, "$"):
			cmd = strings.TrimPrefix(line, "$ ")
			switch {
			case cmd == "cd ..":
				pwd = pwd.p
			case strings.HasPrefix(cmd, "cd"):
				n := strings.TrimPrefix(cmd, "cd ")
				d := &dir{n: n, c: make(map[string]entry), p: pwd}
				pwd.c[n] = d
				pwd = d
			}
		case strings.HasPrefix(line, "dir"):
			n := strings.TrimPrefix(line, "dir ")
			if _, exists := pwd.c[n]; !exists {
				pwd.c[n] = &dir{n: n, c: make(map[string]entry), p: pwd}
			}
		default:
			ss := strings.Fields(line)
			s := exp.Must(strconv.Atoi(ss[0]))
			pwd.c[ss[1]] = &file{n: ss[1], s: s}
		}
	}

	total := 70000000
	need := 30000000
	used := root.size()
	freemin := need - (total - used)

	min := used
	q := []*dir{root}

	for len(q) > 0 {
		d := q[0]
		q = q[1:]
		if s := d.size(); s >= freemin && s < min {
			min = s
		}
		for _, e := range d.c {
			if c, ok := e.(*dir); ok {
				q = append(q, c)
			}
		}
	}
	fmt.Println(min)
}
