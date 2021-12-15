package main

import (
	"bufio"
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
	cave := make([][]int, 5*len(sample))
	for j := range cave {
		cave[j] = make([]int, 5*len(sample[0]))
	}
	for j := range sample {
		for i := range sample[j] {
			cave[j][i] = int(sample[j][i] - '0')
		}
		for i := len(sample[j]); i < len(cave[j]); i++ {
			cave[j][i] = cave[j][i-len(sample[j])] + 1
			for cave[j][i] > 9 {
				cave[j][i] -= 9
			}
		}
	}
	for j := len(sample); j < len(cave); j++ {
		for i := range cave[j] {
			cave[j][i] = cave[j-len(sample)][i] + 1
			for cave[j][i] > 9 {
				cave[j][i] -= 9
			}
		}
	}

	dist := make([][]int, len(cave))
	for i := range dist {
		dist[i] = make([]int, len(cave[0]))
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}
	dist[0][0] = 0

	queue := map[image.Point]struct{}{
		{0, 0}: {},
	}
	for len(queue) > 0 {
		for p := range queue {
			delete(queue, p)
			if p.X > 0 {
				if t := dist[p.Y][p.X] + cave[p.Y][p.X-1]; t < dist[p.Y][p.X-1] {
					dist[p.Y][p.X-1] = t
					queue[image.Point{p.X - 1, p.Y}] = struct{}{}
				}
			}
			if p.X < len(cave[0])-1 {
				if t := dist[p.Y][p.X] + cave[p.Y][p.X+1]; t < dist[p.Y][p.X+1] {
					dist[p.Y][p.X+1] = t
					queue[image.Point{p.X + 1, p.Y}] = struct{}{}
				}
			}
			if p.Y > 0 {
				if t := dist[p.Y][p.X] + cave[p.Y-1][p.X]; t < dist[p.Y-1][p.X] {
					dist[p.Y-1][p.X] = t
					queue[image.Point{p.X, p.Y - 1}] = struct{}{}
				}
			}
			if p.Y < len(cave[0])-1 {
				if t := dist[p.Y][p.X] + cave[p.Y+1][p.X]; t < dist[p.Y+1][p.X] {
					dist[p.Y+1][p.X] = t
					queue[image.Point{p.X, p.Y + 1}] = struct{}{}
				}
			}
			break
		}
	}

	fmt.Println(dist[len(dist)-1][len(dist[len(dist)-1])-1])
}
