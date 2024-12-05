package main

import (
	"fmt"
	"log"
	"math"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

var operation = map[string]func(int, int) int {
	"inc": func(r, n int) int { return r + n },
	"dec": func(r, n int) int { return r - n },
}

var comparison = map[string]func(int, int) bool {
	"<": func(x, y int) bool { return x < y },
	">": func(x, y int) bool { return x > y },
	"<=": func(x, y int) bool { return x <= y },
	">=": func(x, y int) bool { return x >= y },
	"==": func(x, y int) bool { return x == y },
	"!=": func(x, y int) bool { return x != y },
}

func main() {
	maxn := math.MinInt
	regs := make(map[string]int)
	exp.MustForEachLineIn("inputs/8.txt", func(line string) {
		var (
			r, op string
			n int
			q, cmp string
			t int
		)
		if _, err := fmt.Sscanf(line, "%s %s %d if %s %s %d", &r, &op, &n, &q, &cmp, &t); err != nil {
			log.Fatalf("Couldn't scan line: %v", err)
		}
		
		// nb: if either cmp or op is invalid, the corresponding line below would panic naturally
		if comparison[cmp](regs[q], t) {
			regs[r] = operation[op](regs[r], n)
		}
		
		_, mn := algo.MapMax(regs)
		if mn > maxn {
			maxn = mn
		}
	})
	
	fmt.Println(maxn)
}