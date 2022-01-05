package main

import (
	"fmt"
	"math"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	P := intcode.ReadProgram("inputs/7.txt")
	amps := []intcode.VM{P, P, P, P, P}
	best := math.MinInt
	s := []int{0, 1, 2, 3, 4}
	for {
		ch := []chan int{make(chan int)}
		for range s {
			ch = append(ch, make(chan int))
		}
		for i, a := range amps {
			go a.Run(ch[i], ch[i+1])
		}
		for i, n := range s {
			ch[i] <- n
		}
		ch[0] <- 0
		if t := <-ch[5]; t > best {
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
