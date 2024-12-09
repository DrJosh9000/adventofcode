package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"drjosh.dev/exp/algo"
)

//go:embed inputs/9.txt
var input string

type chunk struct {
	id, ln int
}

func dump(start *algo.ListNode[chunk]) {
	p := start
	for {
		if p.Value.id < 0 {
			fmt.Print(strings.Repeat(".", p.Value.ln))
		} else {
			fmt.Print(strings.Repeat(strconv.Itoa(p.Value.id), p.Value.ln))
		}

		p = p.Next
		if p == start {
			break
		}
	}
	fmt.Println()
}

func main() {
	id := 0
	var disksl []chunk
	for i, c := range input {
		if c < '0' || c > '9' {
			break
		}
		switch i % 2 {
		case 0:
			id = i / 2
			disksl = append(disksl, chunk{id, int(c - '0')})
		case 1:
			disksl = append(disksl, chunk{-1, int(c - '0')})
		}
	}

	disk := algo.ListFromSlice(disksl)

	// dump(disk[0])

	for i := range disk {
		j := len(disk) - 1 - i
		if j == 0 {
			break
		}
		ch := disk[j]
		if ch.Value.id < 0 {
			continue
		}

		for p := disk[1]; p != ch && p != disk[0]; p = p.Next {
			if p.Value.id >= 0 {
				continue
			}
			if p.Value.ln < ch.Value.ln {
				continue
			}
			n := &algo.ListNode[chunk]{Value: chunk{id: -1, ln: ch.Value.ln}}
			n.InsertAfter(ch)
			ch.Remove()
			ch.InsertBefore(p)
			p.Value.ln -= ch.Value.ln
			break
		}
		// dump(disk[0])
	}

	// dump(disk[0])

	cksum := 0
	head := 0
	p := disk[0]
	for {
		if p.Value.id < 0 {
			head += p.Value.ln
		} else {
			for range p.Value.ln {
				cksum += head * p.Value.id
				head++
			}
		}

		p = p.Next
		if p == disk[0] {
			break
		}
	}

	fmt.Println(cksum)
}
