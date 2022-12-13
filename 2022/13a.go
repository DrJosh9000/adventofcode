package main

import (
	"encoding/json"
	"fmt"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 13, part a

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
	sum := 0
	index := 1
	var l, r any
	state := 0
	for _, line := range exp.MustReadLines("inputs/13.txt") {
		switch state {
		case 0:
			json.Unmarshal([]byte(line), &l)
			state++
		case 1:
			json.Unmarshal([]byte(line), &r)
			state++

			if cmp(l, r) < 0 {
				sum += index
			}
		case 2:
			index++
			state = 0
		}
	}

	fmt.Println(sum)
}
