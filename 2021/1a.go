package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("1.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var nums []int
	for {
		var n int
		_, err := fmt.Fscan(f, &n)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldn't scan: %v", err)
		}
		nums = append(nums, n)
	}

	count := 0
	prev := nums[0]
	for _, a := range nums[1:] {
		if a > prev {
			count++
		}
		prev = a
	}
	fmt.Println(count)
}
