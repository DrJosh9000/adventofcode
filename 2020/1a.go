package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	var nums []int
	for {
		var n int
		if _, err := fmt.Fscanf(f, "%d", &n); err != nil {
			break
		}
		nums = append(nums, n)
	}

	for i, a := range nums {
		for _, b := range nums[i+1:] {
			if a+b == 2020 {
				fmt.Println(a*b)
				return
			}
		} 
	}
}
