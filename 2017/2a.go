package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
)

func main() {
	sum := 0
	exp.MustForEachLineIn("inputs/2.txt", func(line string) {
		min, max := math.MaxInt, math.MinInt
		for _, ns := range strings.Fields(line) {
			n, err := strconv.Atoi(ns)
			if err != nil {
				log.Fatalf("Couldn't parse %q: %v", ns, err)
			}
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
		sum += max - min
	})
	fmt.Println(sum)
}