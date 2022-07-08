package main

import "fmt"

// been a while since I wrote a circular linked list, so wth let's do it

type node struct {
	value int
	next, prev *node
}

func (n *node) remove() int {
	n.prev.next = n.next
	n.next.prev = n.prev
	return n.value
}

func (n *node) insertAfter(v int) *node {
	o := n.next
	x := &node{
		value: v,
		next: o,
		prev: n,
	}
	o.prev = x
	n.next = x
	return x
}

func main() {
	// can't be bothered with input parsing; just update these numbers
	const players = 9
	const last = 25
	
	cur := &node{value: 0}
	cur.next = cur
	cur.prev = cur
	
	scores := make([]int, players)
	player := 0
	
	for i := 1; i <= last; i++ {
		if i % 23 == 0 {
			scores[player] += i
			cur = cur.prev.prev.prev.prev.prev.prev
			scores[player] += cur.prev.remove()
		} else {
			cur = cur.next.insertAfter(i)
		}
		player++
		player %= players
	}
	
	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	fmt.Println(max)
}