package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Each round of the dance produces a permutation of the items.
// But the p step ("partner") swaps items with given names, not positions,
// so the dance does not produce a single fixed permutation: different starting 
// arrangements lead to different permutations.
// This can be handled by treating p as having the dance partners swap their
// name tags rather than their positions. Thus the dance is a fixed pair of two
// permutations: one permutation of the elements, and another separate
// permutation of the labels.
// Permutation pairs can be composed, and the composition is associative, so we
// can use algo.Pow.

type dance struct {
	line, label [16]byte
}

func (d dance) print() {
	for _, c := range d.line {
		fmt.Printf("%c", d.label[c]+'a')
	}
	fmt.Println()
}

func compose(p, q dance) dance {
	var d dance
	for j, x := range p.line {
		d.line[j] = q.line[x]
	}
	for j, x := range p.label {
		d.label[j] = q.label[x]
	}
	return d
}

func main() {
	var base dance
	var rev [16]byte // label x is at position y
	for i := range base.line {
		base.line[i] = byte(i)
		base.label[i] = byte(i)
		rev[i] = byte(i)
	}
	
	steps := strings.Split(strings.TrimSpace(string(exp.Must(os.ReadFile("inputs/16.txt")))), ",")

	for _, step := range steps {
		switch step[0] {
		case 's':
			var n int 
			exp.Must(fmt.Sscanf(step, "s%d", &n))
			var l [16]byte
			for i := range base.line {
				l[(i+n)%16] = base.line[i]
			}
			base.line = l
		case 'x':
			var a, b int
			exp.Must(fmt.Sscanf(step, "x%d/%d", &a, &b))
			base.line[a], base.line[b] = base.line[b], base.line[a]
		case 'p':
			var a, b rune
			exp.Must(fmt.Sscanf(step, "p%c/%c", &a, &b))
			i, j := rev[a-'a'], rev[b-'a']
			base.label[i], base.label[j] = base.label[j], base.label[i]
			rev[base.label[i]], rev[base.label[j]] = i, j
		}
	}

	//base.print()
	
	algo.Pow(base, 1_000_000_000, compose).print()
}