package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
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
		cs := algo.Keys(lc)
		sort.Slice(cs, func(i, j int) bool {
			if lc[cs[i]] == lc[cs[j]] {
				return cs[i] < cs[j]
			}
			return lc[cs[i]] > lc[cs[j]]
		})
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
