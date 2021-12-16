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
	w, h := 5*len(sample[0]), 5*len(sample)
	cave := make([][]int, h)
	for j := range cave {
		cave[j] = make([]int, w)
	}
	for j := range sample {
		for i := range sample[j] {
			cave[j][i] = int(sample[j][i] - '0')
		}
		for i := len(sample[j]); i < w; i++ {
			cave[j][i] = cave[j][i-len(sample[j])] + 1
			for cave[j][i] > 9 {
				cave[j][i] -= 9
			}
		}
	}
	for j := len(sample); j < h; j++ {
		for i := range cave[j] {
			cave[j][i] = cave[j-len(sample)][i] + 1
			for cave[j][i] > 9 {
				cave[j][i] -= 9
			}
		}
	}

	dist := make([][]int, h)
	for i := range dist {
		dist[i] = make([]int, w)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}
	dist[0][0] = 0

	visited := make(map[image.Point]struct{})
	pq := &priqueue{item{p: image.Pt(0, 0), v: 0}}
	for pq.Len() > 0 {
		it := heap.Pop(pq).(item)
		visited[it.p] = struct{}{}
		for _, d := range []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			p := it.p.Add(d)
			if p.X < 0 || p.X >= w || p.Y < 0 || p.Y >= h {
				continue
			}
			if _, no := visited[p]; no {
				continue
			}
			if t := it.v + cave[p.Y][p.X]; t < dist[p.Y][p.X] {
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
