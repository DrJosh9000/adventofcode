package main

import (
	"fmt"
	"image"
	"sort"

	"drjosh.dev/exp"
)

type dirtrack struct {
	d, t rune
}

var dirs = map[rune]image.Point{
	'>': {1, 0},
	'<': {-1, 0},
	'^': {0, -1},
	'v': {0, 1},
}

var nextDir = map[dirtrack]rune{
	// existing carts = straight track
	{'>', '>'}: '>',
	{'<', '<'}: '<',
	{'^', '^'}: '^',
	{'v', 'v'}: 'v',
	{'>', '<'}: '>',
	{'<', '>'}: '<',
	{'^', 'v'}: '^',
	{'v', '^'}: 'v',
	// straight track
	{'>', '-'}: '>',
	{'<', '-'}: '<',
	{'^', '|'}: '^',
	{'v', '|'}: 'v',
	// crossings
	{'>', '+'}: '>',
	{'<', '+'}: '<',
	{'^', '+'}: '^',
	{'v', '+'}: 'v',
	// turns
	{'>', '\\'}: 'v',
	{'>', '/'}:  '^',
	{'<', '\\'}: '^',
	{'<', '/'}:  'v',
	{'^', '\\'}: '<',
	{'^', '/'}:  '>',
	{'v', '\\'}: '>',
	{'v', '/'}:  '<',
}

type bias int

const (
	left = bias(iota)
	straight
	right
	numBiases
)

type dirbias struct {
	d rune
	b bias
}

var crossings = map[dirbias]rune{
	{'>', left}:     '^',
	{'>', straight}: '>',
	{'>', right}:    'v',
	{'<', left}:     'v',
	{'<', straight}: '<',
	{'<', right}:    '^',
	{'^', left}:     '<',
	{'^', straight}: '^',
	{'^', right}:    '>',
	{'v', left}:     '>',
	{'v', straight}: 'v',
	{'v', right}:    '<',
}

type cart struct {
	image.Point
	d    rune
	b    bias
	dead bool
}

func isCart(x rune) bool {
	return x == '>' || x == '<' || x == '^' || x == 'v'
}

func sortCarts(carts []cart) {
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].dead == carts[j].dead {
			if carts[i].Y == carts[j].Y {
				return carts[i].X < carts[j].X
			}
			return carts[i].Y < carts[j].Y
		}
		return carts[j].dead // => !carts[i].dead
	})
}

func main() {
	var grid []string
	var carts []cart
	r := 0
	exp.MustForEachLineIn("inputs/13.txt", func(line string) {
		grid = append(grid, line)
		for c, x := range line {
			if isCart(x) {
				carts = append(carts, cart{
					Point: image.Pt(c, r),
					d:     x,
				})
			}
		}
		r++
	})

	for tick := 0; ; tick++ {
		sortCarts(carts)
		for i, c := range carts {
			if c.dead {
				carts = carts[:i]
				break
			}
		}
		if len(carts) == 1 {
			fmt.Println(carts[0].Point)
			return
		}
	oneTickLoop:
		for i, c := range carts {
			if c.dead {
				continue
			}
			p := c.Add(dirs[c.d])
			for j, d := range carts {
				if j == i || d.dead {
					continue
				}
				if p == d.Point {
					carts[i].dead = true
					carts[j].dead = true
					continue oneTickLoop
				}
			}
			c.Point = p
			t := rune(grid[p.Y][p.X])
			if t == '+' {
				c.d = crossings[dirbias{c.d, c.b}]
				c.b++
				c.b %= numBiases
			} else {
				c.d = nextDir[dirtrack{c.d, t}]
			}
			carts[i] = c
		}
	}
}
