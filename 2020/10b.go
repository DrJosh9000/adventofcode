
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
	last := nums[len(nums)-1]
	
	// Sounds like a job for DP
	ways := make([]int64, last+1)
	ways[0] = 1
	for _, n := range nums {
		for d:=1; d<4; d++ {
			if n-d < 0 {
				break
			}
			ways[n] += ways[n-d]
		}
	}
	fmt.Println(ways[last])
}
