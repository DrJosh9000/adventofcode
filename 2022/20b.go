package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 20, part b

const key = 811589153
const reps = 10

func main() {
	input := exp.MustReadInts("inputs/20.txt", "\n")

	N := len(input)
	list := algo.ListFromSlice(algo.Map(input, func(x int) image.Point {
		return image.Pt(x*key, (x*key)%(N-1))
	}))
	var zero *algo.ListNode[image.Point]
	for _, n := range list {
		if n.Value.X == 0 {
			zero = n
		}
	}
	for rep := 0; rep < reps; rep++ {
		for _, n := range list {
			if n.Value.Y == 0 {
				continue
			}

			n.Remove()
			p := n.Succ(n.Value.Y)
			if n.Value.Y < 0 {
				n.InsertBefore(p)
			} else {
				n.InsertAfter(p)
			}
		}

		//fmt.Println(zero.ToSlice())
	}

	sum := 0
	p := zero
	for i := 0; i < 3; i++ {
		p = p.Succ(1000)
		fmt.Printf("%d000th number is %d\n", i+1, p.Value.X)
		sum += p.Value.X
	}

	fmt.Println(sum)
}
