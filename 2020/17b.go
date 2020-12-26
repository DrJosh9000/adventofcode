package main

import (
	"bufio"
	"fmt"
	"os"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
	os.Exit(1)
}

type coord struct{ x, y, z, w int }

func (c coord) plus(d coord) coord {
	return coord{c.x + d.x, c.y + d.y, c.z + d.z, c.w + d.w}
}

var neighbors = func() (ns []coord) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					ns = append(ns, coord{i, j, k, l})
				}
			}
		}
	}
	return
}()

var hood = append(neighbors, coord{})

type realm map[coord]struct{}

func cycle(r realm) realm {
	nr, seen := make(realm), make(realm)

	for p := range r {
		for _, n := range hood {
			// Consider the point q: ranges over neighbors of p, plus p itself.
			q := p.plus(n)
			// If we've already computed it, skip.
			if _, s := seen[q]; s {
				continue
			}
			seen[q] = struct{}{}
			// Count the neighbors of q.
			nc := 0
			for _, m := range neighbors {
				if _, active := r[q.plus(m)]; active {
					nc++
				}
			}
			// Apply the rule.
			if _, active := r[q]; (active && nc == 2) || nc == 3 {
				nr[q] = struct{}{}
			}
		}
	}
	return nr
}

func main() {
	f, err := os.Open("input.17")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	row := 0
	realm := make(realm)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		for col, c := range m {
			if c == '#' {
				realm[coord{x: row, y: col}] = struct{}{}
			}
		}
		row++
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	for i := 0; i < 6; i++ {
		realm = cycle(realm)
	}
	fmt.Println(len(realm))
}
