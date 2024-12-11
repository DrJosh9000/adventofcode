package main

import (
	_ "embed"
	"fmt"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

//go:embed inputs/11.txt
var input string

func main() {
	line := algo.Freq(algo.Map(strings.Fields(input), exp.MustAtoi))
	for range 75 {
		line2 := make(map[int]int)
		for n, c := range line {
			if n == 0 {
				line2[1] += c
				continue
			}
			if l := ln10(n); l%2 == 0 {
				left := n
				rf := 1
				right := 0
				for range l / 2 {
					right += rf * (left % 10)
					rf *= 10
					left /= 10
				}
				line2[left] += c
				line2[right] += c
				continue
			}
			line2[n*2024] += c
		}
		line = line2
	}

	sum := 0
	for _, v := range line {
		sum += v
	}
	fmt.Println(sum)
}

func ln10(n int) int {
	c := 0
	for n > 0 {
		n /= 10
		c++
	}
	return c
}
