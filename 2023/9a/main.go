package main

import (
	"fmt"
	"strconv"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/para"
)

// Advent of Code 2023
// Day 9, part a

const inputPath = "2023/inputs/9.txt"

func diffs(in []int) ([]int, bool) {
	allZ := true
	out := make([]int, len(in)-1)
	for i := range out {
		out[i] = in[i+1] - in[i]
		allZ = allZ && out[i] == 0
	}
	return out, allZ
}

func last[S ~[]E, E any](s S) E {
	return s[len(s)-1]
}

func main() {
	lines := exp.MustReadLines(inputPath)
	fmt.Println(algo.Sum(para.Map(lines, func(line string) int {
		nums := [][]int{
			algo.MustMap(strings.Fields(line), strconv.Atoi),
		}
		for {
			ds, allZ := diffs(last(nums))
			nums = append(nums, ds)
			if allZ {
				break
			}
		}
		//fmt.Println(nums)
		for i := range nums[1:] {
			j := len(nums) - i - 1
			nums[j-1] = append(nums[j-1], last(nums[j-1])+last(nums[j]))
		}
		return last(nums[0])
	})))
}
