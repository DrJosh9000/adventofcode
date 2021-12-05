package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/5.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var lines []line
	for {
		var l line
		_, err := fmt.Fscanf(f, "%d,%d -> %d,%d", &l.x1, &l.y1, &l.x2, &l.y2)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldn't scan: %v", err)
		}
		fmt.Println(l)
		// horizontal or vertical
		if l.x1 != l.x2 && l.y1 != l.y2 {
			continue
		}
		lines = append(lines, l)
	}

	points := make(map[point]int)
	for _, l := range lines {
		if l.x1 != l.x2 { // horizontal
			for x, end := minmax(l.x1, l.x2); x <= end; x++ {
				points[point{x, l.y1}]++
			}
		} else { // vertical
			for y, end := minmax(l.y1, l.y2); y <= end; y++ {
				points[point{l.x1, y}]++
			}
		}
	}
	count := 0
	for _, n := range points {
		if n >= 2 {
			count++
		}
	}
	fmt.Println(count)
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

type line struct {
	x1, y1, x2, y2 int
}

type point struct {
	x, y int
}
