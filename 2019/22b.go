package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/22.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()

	const N = 119_315_717_514_047
	const R = 101_741_582_076_661

	// Combine the shuffle instructions into a single linear function.
	// Card x ends up in position y = (mx + b) % N
	// The deck starts in order: m = 1, b = 0.
	m, b := 1, 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		switch {
		case t == "deal into new stack":
			// The deck is reversed; the current function is multiplied by -1
			// and shifted so that 0 becomes N-1.
			// (-1 \equiv N-1 (mod N), so use N-1 to keep it positive.)
			m = mul(m, N-1, N)
			b = mul(b+1, N-1, N)

		case strings.HasPrefix(t, "deal with increment "):
			// The current function is multiplied by n.
			n, err := strconv.Atoi(strings.TrimPrefix(t, "deal with increment "))
			if err != nil {
				log.Fatalf("Invalid increment %q: %v", t, err)
			}
			m = mul(n, m, N)
			b = mul(n, b, N)

		case strings.HasPrefix(t, "cut "):
			// The current function is shifted by n.
			// (Add N to keep it positive.)
			n, err := strconv.Atoi(strings.TrimPrefix(t, "cut "))
			if err != nil {
				log.Fatalf("Invalid cut %q: %v", t, err)
			}
			b = (N + b - n) % N
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan input: %v", err)
	}

	// Let f(x) = (mx + b) % N.
	// Then:
	// f^R(x) = f(f(f(...(x)))) % N
	//        = m(m(m(...m(mx+b)+b...)+b)+b)+b
	//        = m(m(m(...m^2x+mb+b...)+b)+b)+b
	//        = m(m(m(...m^3x+m^2b+mb+b...)+b)+b)+b
	//        = m(m(m(...m^4x+m^3b+m^2b+mb+b...)+b)+b)+b
	//        = ...
	//        = m^Rx + (m^(R-1) + m^(R-2) + ... m + 1)b

	// m^(R-1) + m^(R-2) + ...m^3 + m^2 + m + 1
	//   = (m^R - 1) / (m - 1)     (geometric series)

	// So:
	// f^R(x) = [m^R*x + (m^R-1)*(m-1)^{-1}*b] % N

	mr := pow(m, R, N)
	ms := mul(mr-1, inverse(m-1, N), N)
	br := mul(ms, b, N)

	// f^R(x) = (mr*x + br) % N

	// Finally, for this problem we need the inverse of f^R.
	// f^{-R}(y) = (mr)^{-1} * (y - br) % N
	fmt.Println(mul(inverse(mr, N), N+2020-br, N))
}

func inverse(a, n int) int {
	t, u, r, s := 0, 1, n, a
	for s != 0 {
		q := r / s
		t, u = u, t-q*u
		r, s = s, r-q*s
	}
	if r > 1 {
		panic(strconv.Itoa(a) + " not invertible modulo " + strconv.Itoa(n))
	}
	if t < 0 {
		return t + n
	}
	return t
}

// x*y % n avoiding overflow
func mul(x, y, n int) int {
	x %= n
	y %= n
	s := 0
	for y > 0 {
		if y%2 == 1 {
			s = (s + x) % n
		}
		y /= 2
		x = (2 * x) % n
	}
	return s
}

// x^y % n avoiding overflow
func pow(x, y, n int) int {
	x %= n
	p := 1
	for y > 0 {
		if y%2 == 1 {
			p = mul(p, x, n)
		}
		y /= 2
		x = mul(x, x, n)
	}
	return p
}
