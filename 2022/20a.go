package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 20, part a

type node struct {
	prev, next *node
	x          int
}

func (n *node) remove() {
	// p -> n -> q becomes p -> q
	p, q := n.prev, n.next
	p.next = q
	q.prev = p
}

func (n *node) insertAfter(p *node) {
	q := p.next
	// p -> q becomes p -> n -> q
	p.next = n
	n.prev = p
	n.next = q
	q.prev = n
}

func (n *node) insertBefore(p *node) {
	q := p.prev
	// q -> p becomes q -> n -> p
	q.next = n
	n.prev = q
	n.next = p
	p.prev = n
}

func main() {
	input := exp.MustReadInts("inputs/20.txt", "\n")

	N := len(input)
	list := make([]*node, N)

	var zero *node
	for i, x := range input {
		list[i] = &node{x: x}
		if x == 0 {
			zero = list[i]
		}
	}
	for i, n := range list {
		n.prev = list[(i-1+N)%N]
		n.next = list[(i+1)%N]
	}

	for _, n := range list {
		if n.x == 0 {
			continue
		}
		p := n
		n.remove()
		if n.x < 0 {
			for j := 0; j > n.x; j-- {
				p = p.prev
			}
			n.insertBefore(p)
		} else {
			for j := 0; j < n.x; j++ {
				p = p.next
			}
			n.insertAfter(p)
		}

		// p = list[0]
		// for i := 0; i < N; i++ {
		// 	fmt.Printf("%d ", p.x)
		// 	p = p.next
		// 	if p == list[0] {
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
