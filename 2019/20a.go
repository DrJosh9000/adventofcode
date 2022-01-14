package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math"
	"os"
)

func main() {
	input, err := os.ReadFile("inputs/20.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	maze := bytes.Split(input, []byte{'\n'})
	bounds := image.Rect(0, 0, len(maze[0]), len(maze))

	dp := []image.Point{{1, 0}, {0, 1}}
	dq := [][]image.Point{
		0: {{-1, 0}, {2, 0}},
		1: {{0, -1}, {0, 2}},
	}
	poi := make(map[string][]image.Point)
	rpoi := make(map[image.Point]string)

	for y := range maze {
		for x, b := range maze[y] {
			if !('A' <= b && b <= 'Z') {
				continue
			}
			p := image.Pt(x, y)
			for i, d := range dp {
				pd := p.Add(d)
				if !pd.In(bounds) {
					continue
				}
				c := maze[pd.Y][pd.X]
				if !('A' <= c && c <= 'Z') {
					continue
				}
				label := string([]byte{b, c})
				for _, d := range dq[i] {
					q := p.Add(d)
					if !q.In(bounds) || maze[q.Y][q.X] != '.' {
						continue
					}
					poi[label] = append(poi[label], q)
					rpoi[q] = label
				}
			}
		}
	}

	steps := []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	dist := make([][]int, len(maze))
	for y := range maze {
		dist[y] = make([]int, len(maze[y]))
		for x := range dist[y] {
			dist[y][x] = math.MaxInt
		}
	}
	start, end := poi["AA"][0], poi["ZZ"][0]
	dist[start.Y][start.X] = 0
	p, q := start, []image.Point{start}
	for len(q) > 0 {
		p, q = q[0], q[1:]
		t := dist[p.Y][p.X]
		if p == end {
			fmt.Println(t)
			return
		}

		st := steps
		if label := rpoi[p]; label != "" {
			for _, r := range poi[label] {
				if r != p {
					st = append(st, r.Sub(p))
				}
			}
		}
		for _, d := range st {
			pd := p.Add(d)
			if maze[pd.Y][pd.X] != '.' {
				continue
			}
			if dist[pd.Y][pd.X] <= t+1 {
				continue
			}
			q = append(q, pd)
			dist[pd.Y][pd.X] = t + 1
		}
	}
}
