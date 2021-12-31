package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(os.Args[1], "-")
	lo, err := strconv.Atoi(input[0])
	if err != nil {
		log.Printf("Start of range not a number: %v", err)
	}
	hi, err := strconv.Atoi(input[1])
	if err != nil {
		log.Printf("End of range not a number: %v", err)
	}

	count := 0
numLoop:
	for p := lo; p <= hi; p++ {
		// Going from right to left, the digits never increase.
		adj := false
		d := p % 10
		x := p / 10
		for x > 0 {
			if x%10 > d {
				continue numLoop
			}
			if x%10 == d {
				adj = true
			}
			d = x % 10
			x /= 10
		}
		if adj {
			count++
		}
	}
	fmt.Println(count)
}
