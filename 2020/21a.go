package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	// Each allergen is found in exactly one ingredient.
	// Allergens aren't always marked.
	// When listed, the ingredient that contains each listed allergen *will* be
	// somewhere in the corresponding ingredients list.

	var ingredients [][]string
	a2is := make(map[string]map[string]struct{})
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		// Turns out every line ends with " (contains ...)"
		m = strings.TrimSuffix(m, ")")
		parts := strings.Split(m, " (contains ")
		if len(parts) != 2 {
			die("len(parts) == %d != 2", len(parts))
		}
		ingredsraw := strings.Split(parts[0], " ")
		ingredients = append(ingredients, ingredsraw)
		ingreds := make(map[string]struct{})
		for _, in := range ingredsraw {
			ingreds[in] = struct{}{}
		}
		allergs := strings.Split(parts[1], ", ")
		for _, a := range allergs {
			if a2is[a] == nil {
				// Haven't seen this allergen yet, so it could be any of the
				// ingredients listed.
				a2is[a] = make(map[string]struct{})
				for in := range ingreds {
					a2is[a][in] = struct{}{}
				}
			} else {
				// Seen this allergen before - it must be common between the
				// existing possibilities and the new ones. Remove any missing
				// from the current food's ingredients.
				for in := range a2is[a] {
					if _, found := ingreds[in]; !found {
						delete(a2is[a], in)
					}
				}
			}

		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	// Elimination: if any allergens are known (one possible ingredient)
	// then that ingredient can't be any other allergen (per the rules).
	// ai - allergenic ingredients
	ai := make(map[string]struct{})
	for {
		modified := false
		for a, is := range a2is {
			if len(is) != 1 {
				continue
			}
			for in := range is { // only once
				ai[in] = struct{}{} // known allergenic ingredient
				// eliminate from all other allergen sets
				for b, js := range a2is {
					if b == a {
						continue
					}
					if _, found := js[in]; found {
						modified = true
						delete(js, in)
					}
				}
			}
		}
		if !modified {
			break
		}
	}
	//fmt.Println(a2is)

	// Tally non-allergenic ingredients
	tally := 0
	for _, food := range ingredients {
		for _, in := range food {
			if _, found := ai[in]; !found {
				tally++
			}
		}
	}
	fmt.Println(tally)
}
