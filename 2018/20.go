package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"

	"github.com/DrJosh9000/exp/algo"
)

type set[K comparable] map[K]struct{}

func tokenise(input string) []string {
	var tokens []string
	tstart := 0
	for i, c := range input {
		switch c {
		case '(', '|', ')':
			if i > tstart {
				tokens = append(tokens, input[tstart:i])
			}
			tokens = append(tokens, input[i:i+1])
			tstart = i+1
		}
	}
	if len(input) > tstart {
		tokens = append(tokens, input[tstart:])
	}
	return tokens
}

type expr []any // a sequence of literals and branches

func (e expr) String() string {
	var sb strings.Builder
	for _, x := range e {
		fmt.Fprint(&sb, x)
	}
	return sb.String()
}

type branch []expr // each subexpression in the branch

func (b branch) String() string {
	var sb strings.Builder
	sb.WriteRune('(')
	for i, x := range b {
		if i > 0 {
			sb.WriteRune('|')
		}
		sb.WriteString(x.String())
	}
	sb.WriteRune(')')
	return sb.String()
}

func parseExpr(tokens []string) (expr, int) {
	var e expr
	i := 0
	for i < len(tokens) {
		t := tokens[i]
		switch t {
		case "(":
			i++
			b, n := parseBranch(tokens[i:])
			e = append(e, b)
			i += n + 1
		case "|", ")": // this was one expr inside a branch
			return e, i
		default:
			i++
			e = append(e, t)
		}
	}
	return e, i
}

func parseBranch(tokens []string) (branch, int) {
	var b branch
	i := 0
	for i < len(tokens) {
		t := tokens[i]
		switch t {
		case ")":
			return b, i
		case "|":
			i++
			fallthrough // in case the next alternate is an empty literal
		default: // literal or open-paren
			e, n := parseExpr(tokens[i:])
			b = append(b, e)
			i += n
		}
	}
	return b, i
}

var (
	bits = map[rune]int{
		'N': 0b0001,
		'E': 0b0010,
		'S': 0b0100,
		'W': 0b1000,
	}
	rev = map[rune]rune{
		'N': 'S',
		'E': 'W',
		'S': 'N',
		'W': 'E',
	}
	steps = map[rune]image.Point{
		'N': {0, -1},
		'E': {1, 0},
		'S': {0, 1},
		'W': {-1, 0},
	}
)

func traverse(m map[image.Point]int, c any, p image.Point) set[image.Point] {
	rm := make(set[image.Point])
	switch x := c.(type) {
	case expr:
		rm[p] = struct{}{}
		for _, s := range x {
			nrm := make(set[image.Point])
			for p := range rm {
				for q := range traverse(m, s, p) {
					nrm[q] = struct{}{}
				}
			}
			rm = nrm
		}
		return rm
		
	case branch:
		for _, b := range x {
			for q := range traverse(m, b, p) {
				rm[q] = struct{}{}
			}
		}
		return rm
	
	case string:
		for _, r := range x {
			m[p] |= bits[r]
			p = p.Add(steps[r])
			m[p] |= bits[rev[r]]
		}
		rm[p] = struct{}{}
		
	default:
		log.Fatalf("Encountered a %T in the AST", c)
	}
	return rm
}

func main() {
	inputb, err := os.ReadFile("inputs/20.txt")
	if err != nil {
		log.Fatalf("Couldn't read file: %v", err)
	}
	input := strings.Trim(string(inputb), "^$")
	
	tokens := tokenise(input)
		
	e, _ := parseExpr(tokens)
	if s := e.String(); s != input {
		fmt.Println("got: ", s)
		fmt.Println("want:", input)
		os.Exit(1)
	}
		
	m := map[image.Point]int{}
	traverse(m, e, image.Pt(0, 0))
		
	onek, maxd := 0, 0
	algo.FloodFill(image.Pt(0, 0), func(p image.Point, d int) ([]image.Point, error) {
		if d > maxd {
			maxd = d
		}
		if d >= 1000 {
			onek++
		}
		var next []image.Point
		for c, b := range bits {
			if m[p] & b != 0 {
				next = append(next, p.Add(steps[c]))
			}
		}
		return next, nil
	})
	fmt.Println("Furthest room distance:", maxd)
	fmt.Println("At least 1000 away:", onek)
}
