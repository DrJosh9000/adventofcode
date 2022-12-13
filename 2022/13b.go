package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 13, part b

func cmp(l, r any) int {
	switch l := l.(type) {
	case float64:
		if r, ok := r.(float64); ok {
			switch {
			case l < r:
				return -1
			case l > r:
				return 1
			}
			return 0
		}
		// Promote l to a list
		return cmp([]any{l}, r)
	case []any:
		if r, ok := r.(float64); ok {
			return cmp(l, []any{r})
		}

		r := r.([]any)
		for i := 0; i < len(l) || i < len(r); i++ {
			if i >= len(l) {
				return -1
			}
			if i >= len(r) {
				return 1
			}
			if c := cmp(l[i], r[i]); c != 0 {
				return c
			}
		}
	}
	return 0
}

func main() {
	packets := []any{
		[]any{[]any{2.0}},
		[]any{[]any{6.0}},
	}

	for _, line := range exp.MustReadLines("inputs/13.txt") {
		if line == "" {
			continue
		}
		var a any
		json.Unmarshal([]byte(line), &a)
		packets = append(packets, a)
	}

	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) < 0
	})

	prod := 1
	for i, p := range packets {
		q, ok := p.([]any)
		if !ok || len(q) != 1 {
			continue
		}
		r, ok := q[0].([]any)
		if !ok || len(r) != 1 {
			continue
		}
		if r[0] == 2.0 || r[0] == 6.0 {
			prod *= (i + 1)
		}
	}

	fmt.Println(prod)
}
