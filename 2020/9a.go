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

func main() {
	f, err := os.Open("input.9")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	var nums []int64
	for {
		var num int64
		_, err := fmt.Fscanf(f, "%d\n", &num)
		if err == io.EOF {
			break
		}		
		if err != nil {
			die("Couldn't fscanf: %v", err)
		}
		ln := len(nums)
		if ln < 25 {
			nums = append(nums, num)
			continue
		}
		found := false
searchLoop:
		for i, a := range nums[:24] {
			for _, b := range nums[i+1:] {
				if a != b && num == a+b {
					found = true
					break searchLoop
				}
			}
		}
		if !found {
			fmt.Println(num)
			return
		}
		nums = append(nums[1:], num)
	}
}
