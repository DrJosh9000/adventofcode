package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	serial, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Usage: 11a serialnumber. Error was: %v", err)
	}

	const gridSize = 300
	grid := make([][]int, gridSize+1)
	sums := make([][]int, gridSize+1)
	for y := range grid {
		grid[y] = make([]int, gridSize+1)
		sums[y] = make([]int, gridSize+1)
		if y == 0 {
			continue
		}
		for x := range grid[y] {
			if x == 0 {
				continue
			}
			rackID := x + 10
			grid[y][x] = ((rackID*y+serial)*rackID)/100%10 - 5
			sums[y][x] = grid[y][x] + sums[y][x-1] + sums[y-1][x] - sums[y-1][x-1]
		}
	}

	max := math.MinInt
	var bestX, bestY, bestN int
	for n := 1; n <= gridSize; n++ {
		for x := 1; x <= len(grid[0])-n; x++ {
			for y := 1; y <= len(grid)-n; y++ {
				sum := sums[y+n-1][x+n-1] - sums[y+n-1][x-1] - sums[y-1][x+n-1] + sums[y-1][x-1]
				if sum > max {
					max = sum
					bestX, bestY, bestN = x, y, n
				}
			}
		}
	}
	fmt.Printf("%d,%d,%d\n", bestX, bestY, bestN)
}
