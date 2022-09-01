package main

import (
	"fmt"
	"os"
	"strconv"
	
	"github.com/DrJosh9000/exp"
)

func main() {
	input := exp.Must(strconv.Atoi(os.Args[1]))
	
	length := 1
	pos := 0
	last := -1
	
	for i := 1; i <= 50_000_000; i++ {
		pos += input
		pos %= length
		length++
		pos++
		if pos == 1 {
			last = i
		}
	}
	
	fmt.Println(last)
}