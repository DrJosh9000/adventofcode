package main

import (
	"fmt"
	"math"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	P := intcode.ReadProgram("inputs/7.txt")
	amps := []*intcode.VM{
		{M: P}, {M: P}, {M: P}, {M: P}, {M: P},
	}
	best := math.MinInt
	s := []int{5, 6, 7, 8, 9}
	for {
		for i, n := range s {
			c := make(chan int, 2)
			c <- n
			amps[i].In = c
			if i > 0 {
				amps[i-1].Out = c
			}
		}
		amps[0].In <- 0
		amps[4].Out = make(chan int)
		for _, a := range amps {
			go a.Run()
		}
		var t int
		for t = range amps[4].Out {
			amps[0].In <- t
		}
		if t > best {
			best = t
		}
		if !nextPerm(s) {
			break
		}
	}
	fmt.Println(best)
}

func nextPerm(s []int) bool {
	i := len(s) - 2
	for s[i] > s[i+1] {
		i--
		if i < 0 {
			return false
		}
	}
	j := len(s) - 1
	for s[j] < s[i] {
		j--
	}
	s[i], s[j] = s[j], s[i]
	j = len(s) - 1
	i++
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return true
}
