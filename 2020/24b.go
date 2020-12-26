package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

type coord struct{ x, y int }

func (c coord) add(d coord) coord {
	return coord{c.x + d.x, c.y + d.y}
}

type floor map[coord]struct{}

var neighs = []coord{
	{-1, 1}, {0, 1},
	{-1, 0}, {1, 0},
	{0, -1}, {1, -1},
}

var hood = append(neighs, coord{})

func (f floor) black(c coord) bool {
	_, black := f[c]
	return black
}

func (f floor) flip(c coord) {
	if f.black(c) {
		delete(f, c)
		return
	}
	f[c] = struct{}{}
}

func iterate(f floor) floor {
	nf, q := make(floor), make(floor)
	for c := range f {
		nf[c] = struct{}{}
		for _, h := range hood {
			q[c.add(h)] = struct{}{}
		}
	}

	for c := range q {
		bc := 0
		for _, n := range neighs {
			if f.black(c.add(n)) {
				bc++
			}
		}
		if b := f.black(c); (b && (bc == 0 || bc > 2)) || (!b && bc == 2) {
			nf.flip(c)
		}
	}
	return nf
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	lobby := make(floor)

	//     (-1,1)   (0,1)
	// (-1,0)   (0,0)   (1,0)
	//     (0,-1)   (1,-1)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		var c coord
		for m != "" {
			switch {
			case strings.HasPrefix(m, "e"):
				c.x++
				m = m[1:]
			case strings.HasPrefix(m, "w"):
				c.x--
				m = m[1:]
			case strings.HasPrefix(m, "ne"):
				c.y++
				m = m[2:]
			case strings.HasPrefix(m, "nw"):
				c.x--
				c.y++
				m = m[2:]
			case strings.HasPrefix(m, "se"):
				c.x++
				c.y--
				m = m[2:]
			case strings.HasPrefix(m, "sw"):
				c.y--
				m = m[2:]
			}
		}
		lobby.flip(c)
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	for i := 0; i < 100; i++ {
		lobby = iterate(lobby)
	}
	fmt.Println(len(lobby))
}
