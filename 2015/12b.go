package main

import (
	"encoding/json"
	"fmt"
	"log"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 12, part b

func sum(o any) int {
	s := 0
	switch x := o.(type) {
	case map[string]any:
		for _, y := range x {
			if y == "red" {
				return 0
			}
			s += sum(y)
		}
	case []any:
		for _, y := range x {
			s += sum(y)
		}
	case float64:
		s += int(x)
	case string, bool:
		// skip
	default:
		log.Fatalf("unsupported type: %T", o)
	}
	return s
}

func main() {
	s := 0
	for _, line := range exp.MustReadLines("inputs/12.txt") {
		var o any
		if err := json.Unmarshal([]byte(line), &o); err != nil {
			log.Fatalf("Couldn't unmarshal JSON: %v", err)
		}
		s += sum(o)
	}
	fmt.Println(s)
}
