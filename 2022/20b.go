package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 20, part b

type node struct {
	prev, next *node
	x, y       int
}

func (n *node) remove() {
	// p -> n -> q becomes p -> q
	p, q := n.prev, n.next
	p.next = q
	q.prev = p
}

func (n *node) insertAfter(p *node) {
	if p == n {
		return
	}
	//n.remove()
	q := p.next
	// p -> q becomes p -> n -> q
	p.next = n
	n.prev = p
	n.next = q
	q.prev = n
}

func (n *node) insertBefore(p *node) {
	if p == n {
		return
	}
	//n.remove()
	q := p.prev
	// q -> p becomes q -> n -> p
	q.next = n
	n.prev = q
	n.next = p
	p.prev = n
}

//const key = 811589153

const key = 811589153
const reps = 10

func main() {
	input := exp.MustReadInts("inputs/20.txt", "\n")

	N := len(input)
	list := make([]*node, N)

	var zero *node
	for i, x := range input {
		n := &node{
			x: x * key,
			y: (x * key) % (N - 1),
		}
		list[i] = n
		if x == 0 {
			zero = n
		}
	}
	for i, n := range list {
		n.prev = list[(i-1+N)%N]
		n.next = list[(i+1)%N]
	}

	for rep := 0; rep < reps; rep++ {
		for _, n := range list {
			if n.y == 0 {
				continue
			}
			p := n
			n.remove()
			if n.y < 0 {
				for j := 0; j > n.y; j-- {
					p = p.prev
				}
				n.insertBefore(p)
			} else {
				for j := 0; j < n.y; j++ {
					p = p.next
				}
				n.insertAfter(p)
			}
		}

		// fmt.Printf("after rep %d:\n", rep)
		// p := zero
		// for {
		// 	fmt.Printf("%d ", p.x)
		// 	p = p.next
		// 	if p == zero {
		// 		break
		// 	}
		// }
		// fmt.Println()
	}

	sum := 0
	p := zero
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			p = p.next
		}
		fmt.Printf("%d000th number is %d\n", i+1, p.x)
		sum += p.x
	}

	fmt.Println(sum)
}
