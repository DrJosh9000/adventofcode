package main

import (
	"image"
	"os"
	"fmt"
	"log"
)

func main() {
	f, err := os.Open("inputs/22.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()
	
	var depth int
	var target image.Point
	if _, err := fmt.Fscanf(f, "depth: %d\ntarget: %d,%d\n", &depth, &target.X, &target.Y); err != nil {
		log.Fatalf("Couldn't scan input: %v", err)
	}
	
	fmt.Println("depth:", depth)
	fmt.Println("target:", target)
	
	const (
		mod = 20183
		xmul = 16807
		ymul = 48271
	)
	left, up := image.Pt(-1, 0), image.Pt(0, -1)
	erosion := make(map[image.Point]int)
	var p image.Point
	for p.Y = 0; p.Y <= target.Y; p.Y++ {
		for p.X = 0; p.X <= target.X; p.X++ {
			switch {
			case p.X == 0:
				erosion[p] = (p.Y * ymul + depth) % mod
			case p.Y == 0:
				erosion[p] = (p.X * xmul + depth) % mod
			default:
				erosion[p] = (erosion[p.Add(left)] * erosion[p.Add(up)] + depth) % mod
			}
		}
	}
	erosion[target] = depth % mod
	
	sum := 0
	for _, e := range erosion {
		sum += e % 3
	}
	fmt.Println("risk level:", sum)
}