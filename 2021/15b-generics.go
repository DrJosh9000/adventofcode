package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"image"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("inputs/15.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var sample []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		sample = append(sample, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	// ok just make a copy 25 times larger then
	sw, sh := len(sample[0]), len(sample)
	w, h := 5*sw, 5*sh
	cave := makeGrid[int](w, h)
	for j := range sample {
		for i := range sample[j] {
			cave[j][i] = int(sample[j][i] - '0')
		}
		for i := len(sample[j]); i < w; i++ {
			cave[j][i] = cave[j][i-sw] % 9 + 1
		}
	}
	for j := len(sample); j < h; j++ {
		for i := range cave[j] {
			cave[j][i] = cave[j-sh][i] % 9 + 1
		}
	}

	dist := makeGrid[int](w, h)
	dist.fill(math.MaxInt)
	dist[0][0] = 0

	visited := makeGrid[bool](w, h)
	pq := &priqueue{item{}}
	for pq.Len() > 0 {
		it := heap.Pop(pq).(item)
		if visited[it.p.Y][it.p.X] {
			continue
		}
		visited[it.p.Y][it.p.X] = true
		for _, d := range []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			p := it.p.Add(d)
			if p.X < 0 || p.X >= w || p.Y < 0 || p.Y >= h || visited[p.Y][p.X] {
				continue
			}
			if t := dist[it.p.Y][it.p.X] + cave[p.Y][p.X]; t < dist[p.Y][p.X] {
				dist[p.Y][p.X] = t
				heap.Push(pq, item{p: p, v: t})
			}
		}
	}

	fmt.Println(dist[h-1][w-1])
}

type item struct {
	p image.Point
	v int
}

type priqueue []item

func (pq priqueue) Len() int            { return len(pq) }
func (pq priqueue) Less(i, j int) bool  { return pq[i].v < pq[j].v }
func (pq priqueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priqueue) Push(x interface{}) { *pq = append(*pq, x.(item)) }
func (pq *priqueue) Pop() interface{} {
	n1 := len(*pq) - 1
	i := (*pq)[n1]
	*pq = (*pq)[0:n1]
	return i
}

type grid[T any] [][]T

func makeGrid[T any](w, h int) grid[T] {
	g := make(grid[T], h)
	for j := range g {
		g[j] = make([]T, w)
	}
	return g
}

func (g grid[T]) fill(v T) {
	for _, row := range g {
		for i := range row {
			row[i] = v
		}
	}
}