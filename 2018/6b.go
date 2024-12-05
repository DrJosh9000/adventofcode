package main

import (
	"fmt"
	"image"
	"log"
	"math"

	"drjosh.dev/exp"
)

var step = []image.Point{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func l1(p image.Point) int {
	return abs(p.X) + abs(p.Y)
}

func main() {
	var pts []image.Point
	var bounds image.Rectangle
	bounds.Min = image.Pt(math.MaxInt, math.MaxInt)
	bounds.Max = image.Pt(math.MinInt, math.MinInt)
	exp.MustForEachLineIn("inputs/6.txt", func(line string) {
		var p image.Point
		if _, err := fmt.Sscanf(line, "%d, %d", &p.X, &p.Y); err != nil {
			log.Fatalf("Couldn't parse line %q: %v", line, err)
		}
		pts = append(pts, p)
		if p.X < bounds.Min.X {
			bounds.Min.X = p.X
		}
		if p.X > bounds.Max.X {
			bounds.Max.X = p.X
		}
		if p.Y < bounds.Min.Y {
			bounds.Min.Y = p.Y
		}
		if p.Y > bounds.Max.Y {
			bounds.Max.Y = p.Y
		}
	})

	count := 0
	var p image.Point
	for p.X = bounds.Min.X; p.X <= bounds.Max.X; p.X++ {
		for p.Y = bounds.Min.Y; p.Y <= bounds.Max.Y; p.Y++ {
			sum := 0
			for _, q := range pts {
				sum += l1(p.Sub(q))
			}
			if sum < 10000 {
				count++
			}
		}
	}

	fmt.Println(count)
}
