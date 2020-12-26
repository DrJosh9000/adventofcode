package main

import (
	"fmt"
	"io"
	"os"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func summinmax(n []int64) int64 {
	min := int64(0x7fffffffffffffff)
	max := int64(-1)
	for _, x := range n {
		if x < min { min = x }
		if x > max { max = x }
	}
	return min + max
}

func main() {
	f, err := os.Open("input.9")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	const targ = 393911906
	var sum int64
	var nums, sums []int64
	for {
		var num int64
		_, err := fmt.Fscanf(f, "%d\n", &num)
		if err == io.EOF {
			break
		}		
		if err != nil {
			die("Couldn't fscanf: %v", err)
		}
		sum += num
		nums = append(nums, num)
		if sum == targ {
			fmt.Println(summinmax(nums))
			return
		}
		for i, x := range sums {
			if sum - x == targ {
				fmt.Println(summinmax(nums[i+1:]))
				return
			}
		}
		sums = append(sums, sum)
	}
}
