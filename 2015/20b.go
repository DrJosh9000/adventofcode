package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2015
// Day 20, part b

func main() {
	target := exp.Must(strconv.Atoi(os.Args[1]))

	var h []int
	for i := 1; ; i++ {
		if len(h) <= 50*i {
			h = append(h, make([]int, 50*i-len(h)+1)...)
		}
		for j := 1; j <= 50; j++ {
			h[i*j] += 11 * i
		}
		if h[i] >= target {
			fmt.Println(i)
			return
		}
	}
}
