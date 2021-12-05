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
		lines = append(lines, l)
	}

	points := make(map[point]int)
	for _, l := range lines {
		xstep, ystep := sign(l.x2-l.x1), sign(l.y2-l.y1)
		for x, y := l.x1, l.y1; x != l.x2+xstep || y != l.y2+ystep; x, y = x+xstep, y+ystep {
			points[point{x, y}]++
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

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

type line struct {
	x1, y1, x2, y2 int
}

type point struct {
	x, y int
}
