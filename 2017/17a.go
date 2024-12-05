package main

import (
	"fmt"
	"os"
	"strconv"

	"drjosh.dev/exp"
)

type node struct {
	value int
	next  *node
}

func main() {
	input := exp.Must(strconv.Atoi(os.Args[1]))

	cur := &node{value: 0}
	cur.next = cur

	for i := 1; i <= 2017; i++ {
		for j := 0; j < input; j++ {
			cur = cur.next
		}
		cur.next = &node{
			value: i,
			next:  cur.next,
		}
		cur = cur.next
	}

	fmt.Println(cur.next.value)
}
