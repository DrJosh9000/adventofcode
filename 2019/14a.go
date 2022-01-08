package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/14.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()

	rules := make(map[string]recipe)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		spl := strings.SplitN(sc.Text(), " => ", 2)
		var ins []item
		for _, it := range strings.Split(spl[0], ", ") {
			ins = append(ins, mustParseItem(it))
		}
		out := mustParseItem(spl[1])
		rules[out.chem] = recipe{
			out: out,
			in:  ins,
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}

	q := []item{{"FUEL", 1}}
	inv := make(map[string]int)
	var it item
	ore := 0
	for len(q) > 0 {
		it, q = q[0], q[1:]
		if it.chem == "ORE" {
			ore += it.amt
			continue
		}
		// Satisfy some demand from inventory
		if stock := inv[it.chem]; stock > 0 {
			cons := min(it.amt, stock)
			it.amt -= cons
			inv[it.chem] -= cons
			if it.amt == 0 {
				continue
			}
		}
		// Create some more based on the recipe
		recipe := rules[it.chem]
		mult := ((it.amt - 1) / recipe.out.amt) + 1
		for _, in := range recipe.in {
			q = append(q, item{in.chem, mult * in.amt})
		}
		inv[it.chem] += mult*recipe.out.amt - it.amt
	}
	fmt.Println(ore)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type recipe struct {
	out item
	in  []item
}

type item struct {
	chem string
	amt  int
}

func mustParseItem(s string) item {
	var itm item
	if _, err := fmt.Sscanf(s, "%d %s", &itm.amt, &itm.chem); err != nil {
		log.Fatalf("Couldn't scan item: %v", err)
	}
	return itm
}
