package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strings"
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
	best := math.MaxInt
	s := state{p: poi['@']}
	dist := map[state]int{s: 0}
	pq := &priqueue{s}
	for pq.Len() > 0 {
		s = heap.Pop(pq).(state)
		s.d = 0
		if s.k == goal && dist[s] < best {
			best = dist[s]
			continue
		}
		for _, ns := range fill(maze, s) {
			nd := ns.d + dist[s]
			ns.d = 0
			if od, seen := dist[ns]; seen && nd >= od {
				continue
			}
			dist[ns] = nd
			heap.Push(pq, ns)
		}
	}
	fmt.Println(best)
}

type state struct {
	p image.Point
	k uint
	d int
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

type priqueue []state

func (pq priqueue) Len() int            { return len(pq) }
func (pq priqueue) Less(i, j int) bool  { return pq[i].d < pq[j].d }
func (pq priqueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priqueue) Push(x interface{}) { *pq = append(*pq, x.(state)) }
func (pq *priqueue) Pop() interface{} {
	n1 := len(*pq) - 1
	i := (*pq)[n1]
	*pq = (*pq)[0:n1]
	return i
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

func fill(maze [][]byte, s state) []state {
	bounds := image.Rect(0, 0, len(maze[0]), len(maze))
	dist := makeGrid(len(maze), len(maze[0]), math.MaxInt)
	dist[s.p.Y][s.p.X] = 0
	var out []state
	p, q := image.Point{}, []image.Point{s.p}
	for len(q) > 0 {
		p, q = q[0], q[1:]
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
			if 'a' <= b && b <= 'z' {
				if k := uint(1) << (b - 'a'); s.k&k == 0 {
					out = append(out, state{
						p: t,
						k: s.k | k,
						d: dist[t.Y][t.X],
					})
				}
			}
		}
	}
	return out
}
