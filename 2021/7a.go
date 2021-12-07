package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inputs/7.txt")
	if err != nil {
		log.Fatalf("Couldn't read: %v", err)
	}

	var crabs []int
	min, max := math.MaxInt, math.MinInt
	for _, s := range strings.Split(string(f), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Couldn't atoi: %v", err)
		}
		crabs = append(crabs, n)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	minfuel := math.MaxInt
	for p := min; p <= max; p++ {
		fuel := 0
		for _, n := range crabs {
			fuel += abs(p - n)
		}
		if fuel < minfuel {
			minfuel = fuel
		}
	}

	fmt.Println(minfuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
