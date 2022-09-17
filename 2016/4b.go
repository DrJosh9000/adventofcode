package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"golang.org/x/exp/maps"
)

// Advent of Code 2016
// Day 4, part b

func main() {
	for _, line := range exp.MustReadLines("inputs/4.txt") {
		lc := make(map[rune]int)
		parts := strings.FieldsFunc(line, func(c rune) bool {
			return c == '-' || c == '[' || c == ']'
		})
		for _, part := range parts[:len(parts)-2] {
			for _, c := range part {
				lc[c]++
			}
		}
		cs := maps.Keys(lc)
		algo.SortAsc(cs)
		algo.SortByMapDesc(cs, lc)
		if string(cs[:5]) != parts[len(parts)-1] {
			// decoy
			continue
		}
		sector := exp.Must(strconv.Atoi(parts[len(parts)-2]))
		parts = parts[:len(parts)-2]
		var sb strings.Builder
		for _, part := range parts {
			for _, c := range part {
				sb.WriteRune('a' + (c-'a'+rune(sector))%26)
			}
			sb.WriteRune(' ')
		}
		if sb.String() == "northpole object storage " {
			fmt.Println(sector)
			return
		}
	}
}
