package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strings"

	"github.com/DrJosh9000/exp/algo"
)

func main() {
	f, err := os.ReadFile("inputs/18.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	maze := bytes.Split(f, []byte{'\n'})

	poi := make(map[byte]image.Point)
	for y, row := range maze {
		for x, b := range row {
			if b != '.' && b != '#' {
				poi[b] = image.Pt(x, y)
			}
		}
	}

	var goal uint
	for b := range poi {
		if 'a' <= b && b <= 'z' {
			goal |= 1 << (b - 'a')
		}
	}
	s := state{p: poi['@']}

	visit := func(s state, d int) ([]algo.WeightedItem[state, int], error) {
		if s.k == goal {
			fmt.Println(d)
			return nil, fmt.Errorf("all done")
		}
		return fill(maze, s), nil
	}

	algo.Dijkstra(s, visit)
}

type state struct {
	p image.Point
	k uint
}

func (s state) String() string {
	var sb strings.Builder
	sb.WriteString(s.p.String())
	sb.WriteByte(' ')
	for i := uint(0); i < 26; i++ {
		if s.k&(1<<i) != 0 {
			sb.WriteByte(byte(i + 'a'))
		}
	}
	return sb.String()
}

var steps = []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func makeGrid(h, w, v int) [][]int {
	out := make([][]int, h)
	for j := range out {
		out[j] = make([]int, w)
		for i := range out[j] {
			out[j][i] = v
		}
	}
	return out
}

func fill(maze [][]byte, s state) []algo.WeightedItem[state, int] {
	bounds := image.Rect(0, 0, len(maze[0]), len(maze))
	dist := makeGrid(len(maze), len(maze[0]), math.MaxInt)
	dist[s.p.Y][s.p.X] = 0
	var out []algo.WeightedItem[state, int]
	p, q := image.Point{}, []image.Point{s.p}
	for len(q) > 0 {
		p, q = q[0], q[1:]

		if b := maze[p.Y][p.X]; 'a' <= b && b <= 'z' {
			if k := uint(1)<<(b-'a'); s.k&k == 0 {
				out = append(out, algo.WeightedItem[state, int]{
					Item: state{
						p: p,
						k: s.k | k,
					},
					Weight: dist[p.Y][p.X],
				})
			}
		}

		for _, d := range steps {
			t := p.Add(d)
			if !t.In(bounds) {
				continue
			}
			if dist[t.Y][t.X] <= dist[p.Y][p.X]+1 {
				continue
			}
			b := maze[t.Y][t.X]
			if b == '#' {
				continue
			}
			if 'A' <= b && b <= 'Z' {
				if s.k&(1<<(b-'A')) == 0 {
					// don't have this key yet
					continue
				}
			}
			q = append(q, t)
			dist[t.Y][t.X] = dist[p.Y][p.X] + 1
		}
	}
	return out
}
