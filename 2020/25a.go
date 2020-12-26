package main

// Great. Our task is to break Diffie-Hellman.

// 1327981
// 2822615

const modulus = 20201227

func pow(b, e int) int {
	switch e {
	case 0:
		return 1
	case 1:
		return b
	}
	h := pow(b, e/2)
	k := 1
	if e%2 == 1 {
		k = b
	}
	return (((h * h) % modulus) * k) % modulus
}

func log7(p int) int {
	c, x := 0, 1
	for ; x != p; c++ {
		x *= 7
		x %= modulus
	}
	return c
}

func main() {
	a, b := 1327981, 2822615
	secret := log7(a)
	println(pow(b, secret))
}
