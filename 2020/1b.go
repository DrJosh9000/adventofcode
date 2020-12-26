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

	for i, a := range nums[:len(nums)-2] {
		for j, b := range nums[i+1:len(nums)-1] {
			for _, c := range nums[i+j+2:] {
				if a+b+c == 2020 {
					fmt.Println(a*b*c)
					return
				}
			}
		} 
	}
}
