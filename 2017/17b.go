package main

import (
	"fmt"
	"os"
	"strconv"
	
	"github.com/DrJosh9000/exp"
)

// Took several minutes but it worked

type node struct {
	value int
	next *node
}

func main() {
	input := exp.Must(strconv.Atoi(os.Args[1]))
	
	zero := &node{value: 0}
	zero.next = zero
	cur := zero
	
	for i := 1; i <= 50_000_000; i++ {
		for j := 0; j < input; j++ {
			cur = cur.next
		}
		cur.next = &node{
			value: i,
			next: cur.next,
		}
		cur = cur.next
	}
	
	fmt.Println(zero.next.value)
}