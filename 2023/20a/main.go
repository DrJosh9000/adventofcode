package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2023
// Day 20, part a

const inputPath = "2023/inputs/20.txt"

type signal bool

const (
	high = signal(false)
	low  = signal(true)
)

func (s signal) String() string {
	if s == high {
		return "high"
	}
	return "low"
}

type module interface {
	outs() []string
	process(signal, string) (signal, bool)
}

type broadcaster struct {
	out []string
}

func (m *broadcaster) process(s signal, _ string) (signal, bool) { return s, true }

type flipflop struct {
	out   []string
	state bool
}

func (m *flipflop) process(s signal, _ string) (signal, bool) {
	if s == high {
		return s, false
	}
	m.state = !m.state
	if m.state {
		return high, true
	}
	return low, true
}

type conjunction struct {
	out  []string
	last map[string]signal
}

func (m *conjunction) process(s signal, src string) (signal, bool) {
	m.last[src] = s
	for _, v := range m.last {
		if v == low {
			return high, true
		}
	}
	return low, true
}

func (m *broadcaster) outs() []string { return m.out }
func (m *flipflop) outs() []string    { return m.out }
func (m *conjunction) outs() []string { return m.out }

func main() {
	lines := exp.MustReadLines(inputPath)
	byName := make(map[string]module)
	for _, line := range lines {
		in, out, ok := strings.Cut(line, " -> ")
		if !ok {
			panic("invalid line " + line)
		}
		outs := strings.Split(out, ", ")
		var m module
		var name string
		switch {
		case in[0] == '%':
			name = in[1:]
			m = &flipflop{
				out:   outs,
				state: false,
			}

		case in[0] == '&':
			name = in[1:]
			m = &conjunction{
				out:  outs,
				last: make(map[string]signal),
			}

		case in == "broadcaster":
			name = in
			m = &broadcaster{
				out: outs,
			}

		default:
			panic(fmt.Sprintf("invalid name %q", name))
		}

		//fmt.Printf("name %q, module %v\n", name, m)
		byName[name] = m
	}

	// conjunctions initially remember low for all inputs
	for name, mod := range byName {
		for _, out := range mod.outs() {
			dmod := byName[out]
			if conj, ok := dmod.(*conjunction); ok {
				conj.last[name] = low
			}
		}
	}

	type pulse struct {
		src, dst string
		sig      signal
	}
	highs, lows := 0, 0
	for i := 0; i < 1000; i++ {
		q := []pulse{{src: "button", dst: "broadcaster", sig: low}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.sig == high {
				highs++
			} else {
				lows++
			}
			dmod := byName[p.dst]
			if dmod == nil {
				//panic("no module named " + p.dst)
				continue
			}
			ns, ok := dmod.process(p.sig, p.src)
			if !ok {
				continue
			}
			for _, out := range dmod.outs() {
				q = append(q, pulse{
					src: p.dst,
					dst: out,
					sig: ns,
				})
			}
		}
	}

	fmt.Println(highs * lows)
}
