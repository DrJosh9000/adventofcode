
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.10")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	var nums []int
	for {
		var num int
		_, err := fmt.Fscanf(f, "%d\n", &num)
		if err == io.EOF {
			break
		}		
		if err != nil {
			die("Couldn't fscanf: %v", err)
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	diffs := make([]int, 4)
	diffs[nums[0]]++
	for i, n := range nums[1:] {
		diffs[n-nums[i]]++
	}
	diffs[3]++ // last adapter diff always 3
	fmt.Println(diffs)
	fmt.Println(diffs[1]*diffs[3])
}
