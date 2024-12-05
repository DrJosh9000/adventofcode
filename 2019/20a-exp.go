package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math"
	"os"

	"drjosh.dev/exp/algo"
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

	type node struct {
		label string
		inner bool
	}

	poi := make(map[node]image.Point)
	rpoi := make(map[image.Point]node)

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
				for _, d := range dq[i] {
					q := p.Add(d)
					if !q.In(bounds) || maze[q.Y][q.X] != '.' {
						continue
					}
					outer := (q.X == 2 || q.Y == 2 || q.X == len(maze[y])-3 || q.Y == len(maze)-3)
					n := node{string([]byte{b, c}), !outer}
					poi[n] = q
					rpoi[q] = n
				}
			}
		}
	}

	steps := []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	graph := make(map[node][]algo.WeightedItem[node, int])
	for sn := range poi {
		if sn.label == "AA" || sn.label == "ZZ" {
			continue
		}
		en := sn
		en.inner = !en.inner
		graph[sn] = append(graph[sn], algo.WeightedItem[node, int]{
			Item:   en,
			Weight: 1,
		})
	}

	for sn, start := range poi {
		dist := make([][]int, len(maze))
		for y := range maze {
			dist[y] = make([]int, len(maze[y]))
			for x := range dist[y] {
				dist[y][x] = math.MaxInt
			}
		}
		dist[start.Y][start.X] = 0
		p, q := start, []image.Point{start}
		for len(q) > 0 {
			p, q = q[0], q[1:]
			t := dist[p.Y][p.X]
			if en, ok := rpoi[p]; ok && p != start {
				graph[sn] = append(graph[sn], algo.WeightedItem[node, int]{
					Item:   en,
					Weight: t,
				})
			}

			for _, d := range steps {
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

	sn, en := node{"AA", false}, node{"ZZ", false}
	algo.Dijkstra(sn, func(n node, d int) ([]algo.WeightedItem[node, int], error) {
		if n == en {
			fmt.Println(d)
			return nil, fmt.Errorf("all done")
		}
		return graph[n], nil
	})
}
