package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.13")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	var unused int
	var rawids string
	var n1 *big.Int
	sol := big.NewInt(0)
	fmt.Fscanf(f, "%d\n%s\n", &unused, &rawids)
	for i, t := range strings.Split(rawids, ",") {
		if t == "x" {
			continue
		}
		id, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			die("Couldn't atoi %q: %v", t, err)
		}

		n2 := big.NewInt(id)
		if n1 == nil {
			n1 = n2
			continue
		}
		// Find bezout coefficients m1, m2
		m1, m2 := big.NewInt(0), big.NewInt(0)
		big.NewInt(0).GCD(m1, m2, n1, n2)
		a1 := big.NewInt(0)
		a1.Set(sol)
		a1.Mul(a1, m2)
		a1.Mul(a1, n2)
		a2 := big.NewInt(id - int64(i))
		a2.Mul(a2, m1)
		a2.Mul(a2, n1)
		sol.Add(a1, a2)

		// n1 is the cumulative product of ids
		n1.Mul(n1, n2)
	}
	fmt.Println(sol.Mod(sol, n1))
}
