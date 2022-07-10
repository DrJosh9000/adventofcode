package main

import (
	"fmt"
	"image"
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

	grid := make([][]int, 300)
	for j := range grid {
		grid[j] = make([]int, 300)
		for i := range grid[j] {
			x, y := i+1, j+1
			rackID := x + 10
			grid[j][i] = ((rackID*y+serial)*rackID)/100%10 - 5
		}
	}

	// fmt.Println(grid[4][2])
	// fmt.Println(grid[78][121])
	// fmt.Println(grid[195][216])
	// fmt.Println(grid[152][100])
	
	max := math.MinInt
	var best image.Point
	for x := 1; x <= len(grid[0])-2; x++ {
		for y := 1; y <= len(grid)-2; y++ {
			sum := 0
			for j := -1; j <= 1; j++ {
				for i := -1; i <= 1; i++ {
					sum += grid[y+j][x+i]
				}
			}
			if sum > max {
				max = sum
				best = image.Pt(x, y)
			}
		}
	} 
	fmt.Println(best)
}
