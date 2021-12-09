package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("inputs/9.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var heightmap []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		heightmap = append(heightmap, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	down := make(map[point]point)
	for y, row := range heightmap {
		for x := range row {
			cell := row[x]
			if cell == '9' {
				continue
			}
			down[point{x, y}] = point{x, y}
			if x > 0 && cell > row[x-1] {
				down[point{x, y}] = point{x - 1, y}
			}
			if x < len(row)-1 && cell > row[x+1] {
				down[point{x, y}] = point{x + 1, y}
			}
			if y > 0 && cell > heightmap[y-1][x] {
				down[point{x, y}] = point{x, y - 1}
			}
			if y < len(heightmap)-1 && cell > heightmap[y+1][x] {
				down[point{x, y}] = point{x, y + 1}
			}
		}
	}

	sizes := make(map[point]int)
	for p := range down {
		for p != down[p] {
			down[p] = down[down[p]]
			p = down[p]
		}
		sizes[p]++
	}
	var all []int
	for _, s := range sizes {
		all = append(all, s)
	}
	sort.Ints(all)
	n := len(all)
	fmt.Println(all[n-1] * all[n-2] * all[n-3])
}

type point struct{ x, y int }
