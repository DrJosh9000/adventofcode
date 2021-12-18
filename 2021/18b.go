package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type any = interface{}

func main() {
	f, err := os.Open("inputs/18.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var numbers [][]any
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var n []any
		if err := json.Unmarshal([]byte(sc.Text()), &n); err != nil {
			log.Fatalf("Couldn't parse line: %v", err)
		}
		numbers = append(numbers, intise(n).([]any))
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	maxmag := 0
	for i, x := range numbers {
		for j, y := range numbers {
			if i == j {
				continue
			}
			if m := magnitude(reduce([]any{clone(x), clone(y)})); m > maxmag {
				maxmag = m
			}
		}
	}

	fmt.Println(maxmag)
}

func clone(x any) any {
	switch x := x.(type) {
	case int:
		return int(x)
	case []any:
		return []any{clone(x[0]), clone(x[1])}
	}
	panic("invalid type in expression")
}

// intise converts all float64s (produced by encoding/json) into ints
func intise(x any) any {
	switch x := x.(type) {
	case float64:
		return int(x)
	case []any:
		return []any{intise(x[0]), intise(x[1])}
	}
	panic("invalid type in expression")
}

// magnitude computes the magnitude of a pair or int.
func magnitude(x any) int {
	switch x := x.(type) {
	case int:
		return x
	case []any:
		return 3*magnitude(x[0]) + 2*magnitude(x[1])
	}
	panic("invalid type in expression")
}

// reduce applies reductions (explode or split).
func reduce(x []any) []any {
	for {
		if explode(x) {
			continue
		}
		if s, did := split(x); did {
			x = s.([]any)
			continue
		}
		break
	}
	return x
}

// deepest returns the path to the leftmost deepest pair.
func deepest(x []any, prefix []int) []int {
	var lp, rp []int
	if l, ok := x[0].([]any); ok {
		lp = deepest(l, []int{0})
	}
	if r, ok := x[1].([]any); ok {
		rp = deepest(r, []int{1})
	}
	if lp == nil && rp == nil {
		return prefix
	}
	if len(rp) > len(lp) {
		return append(prefix, rp...)
	}
	return append(prefix, lp...)
}

// explode explodes the first pair nested at least 4 deep.
func explode(x []any) bool {
	path := deepest(x, nil)
	if len(path) < 4 {
		return false
	}
	// find the target pair and the path of pairs
	var stack [][]any
	targ := x
	for _, t := range path {
		stack = append(stack, targ)
		targ = targ[t].([]any)
	}
	l, r := targ[0].(int), targ[1].(int)
	// find the rightmost number to the left of the path, add l to it
	// first walk up the path to find the most recent pair where we could have
	// gone left instead of right...
	for i := range path {
		j := len(path) - i - 1
		if path[j] == 0 {
			continue
		}
		// then proceed rightward down the left branch of the subtree from here
		par := stack[j]
		if n, yes := par[0].(int); yes {
			par[0] = n + l
			break
		}
		par = par[0].([]any)
		for {
			if n, yes := par[1].(int); yes {
				par[1] = n + l
				break
			} else {
				par = par[1].([]any)
			}
		}
		break
	}

	// find the leftmost number to the right of the path, add r to it
	for i := range path {
		j := len(path) - i - 1
		if path[j] == 1 {
			continue
		}
		par := stack[j]
		if n, yes := par[1].(int); yes {
			par[1] = n + r
			break
		}
		par = par[1].([]any)
		for {
			if n, yes := par[0].(int); yes {
				par[0] = n + r
				break
			} else {
				par = par[0].([]any)
			}
		}
		break
	}

	// finally, replace the target with 0
	stack[len(stack)-1][path[len(path)-1]] = 0
	return true
}

// split splits the leftmost number larger than 9
func split(x any) (any, bool) {
	switch x := x.(type) {
	case int:
		if x > 9 {
			return []any{x / 2, x - (x / 2)}, true
		}
		return x, false
	case []any:
		if l, found := split(x[0]); found {
			return []any{l, x[1]}, true
		}
		r, found := split(x[1])
		return []any{x[0], r}, found
	}
	panic("invalid type in expression")
}
