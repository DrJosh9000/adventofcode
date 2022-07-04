package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	in, err := os.ReadFile("inputs/5.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}
	
	var reps []string
	for c := 'A'; c <= 'Z'; c++ {
		reps = append(reps, 
			fmt.Sprintf("%c%c", c, c+32), "",
			fmt.Sprintf("%c%c", c+32, c), "",
		)
	}
	
	r := strings.NewReplacer(reps...)
	reduce := func(s string) int {
		for {
			n := r.Replace(s)
			if n == s {
				return len(s)
			}
			s = n
		}
	}
	
	s := strings.TrimSpace(string(in))
	min := math.MaxInt
	for c := 'A'; c <= 'Z'; c++ {
		t := strings.Map(func(r rune) rune {
			if r == c || r == c+32 { return -1 }
			return r
		}, s)
		if l := reduce(t); l < min {
			min = l
		}
	}
	
	fmt.Println(min)
}