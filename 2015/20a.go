package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2015
// Day 20, part a

func main() {
	target := exp.Must(strconv.Atoi(os.Args[1])) / 10

	h := make([]int, target+1)
	for i := 1; i <= target; i++ {
		for j := i; j <= target; j += i {
			h[j] += i
		}
		if h[i] >= target {
			fmt.Println(i)
			return
		}
	}
}
