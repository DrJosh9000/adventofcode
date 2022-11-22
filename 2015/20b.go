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

	type house struct {
		n, p int
	}
	h := make(map[int]int)
	ch := make(chan house)
	go func() {
		for i := 1; ; i++ {
			for j := 1; j <= 50; j++ {
				h[i*j] += 11 * i
			}
			ch <- house{i, h[i]}
		}
	}()

	for h := range ch {
		if h.p >= target {
			fmt.Println(h.n)
			return
		}
	}
}
