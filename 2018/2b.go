package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

func main() {
	ids := []string{}
	exp.MustForEachLineIn("inputs/2.txt", func(line string) {
		ids = append(ids, line)
	})
	for i, a := range ids[:len(ids)-1] {
	idsLoop:
		for _, b := range ids[i+1:] {
			d := 0
			p := -1
			for j := range a {
				if a[j] == b[j] {
					continue
				}
				d++
				if d > 1 {
					continue idsLoop
				}
				p = j
			}
			if d != 1 {
				continue
			}
			for j, r := range a {
				if j == p {
					continue
				}
				fmt.Printf("%c", r)
			}
			fmt.Println()
			return
		}
	}
}