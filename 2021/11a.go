package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/11.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var grid [][]byte
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		row := make([]byte, 0, len(line))
		for _, r := range line {
			row = append(row, byte(r-'0'))
		}
		grid = append(grid, row)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	flashes := 0
	for i := 0; i < 100; i++ {
		var queue []point
		seen := make(map[point]struct{})
		for y, row := range grid {
			for x := range row {
				grid[y][x]++
			}
		}
		for y, row := range grid {
			for x := range row {
				if grid[y][x] > 9 {
					p := point{x, y}
					queue = append(queue, p)
					seen[p] = struct{}{}
				}
			}
		}
		for len(queue) > 0 {
			q := queue[0]
			queue = queue[1:]
			for _, d := range surrounds {
				p := point{d.x + q.x, d.y + q.y}
				if p.x >= 0 && p.x < 10 && p.y >= 0 && p.y < 10 {
					grid[p.y][p.x]++
					if grid[p.y][p.x] > 9 {
						if _, no := seen[p]; !no {
							queue = append(queue, p)
							seen[p] = struct{}{}
						}
					}
				}
			}
		}
		for y, row := range grid {
			for x := range row {
				if grid[y][x] > 9 {
					grid[y][x] = 0
					flashes++
				}
			}
		}
	}
	fmt.Println(flashes)
}

type point struct{ x, y int }

var surrounds = []point{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
}

func copyGrid(g [][]byte) [][]byte {
	n := make([][]byte, len(g))
	for i, r := range g {
		n[i] = make([]byte, len(r))
		copy(n[i], r)
	}
	return n
}
