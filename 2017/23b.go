package main

import "fmt"

/*
	b = 93
	c = b
	if a != 0 { goto l1 }
	goto l2
l1: b *= 100
	b += 100000
	c = b
	c += 17000
l2: f = 1
	d = 2
l5: e = 2
l4: g = d
	g *= e
	g -= b
	if g != 0 { goto l3 }
	f = 0
l3: e++
	g = e
	g -= b
	if g != 0 { goto l4 }
	d++
	g = d
	g -= b
	if g != 0 { goto l5 }
	if f != 0 { goto l6 }
	h++
l6: g = b
	g -= c
	if g != 0 { goto l7 }
	goto l8
l7: b += 17
	goto l2
l8:

-------

b = 93
c = b
if a != 0 {
	b *= 100
	b += 100000
	c = b+17000
}

for {
	f = 1
	d = 2
	for {
		e = 2
		for {
			if d*e - b == 0 {
				f = 0
			}
			e++
			if e-b == 0 {
				break
			}
		}
		d++
		if d-b == 0 {
			break
		}
	}
	if f == 0 {
		h++
	}
	if b-c == 0 {
		break
	}
	b += 17
}


---

a = 1
b = 109300
c = 126300
for {
	f = 1
	d = 2
	for {
		e = 2
		for {
			if d*e == b {
				f = 0
			}
			e++
			if e == b {
				break
			}
		}
		d++
		if d == b {
			break
		}
	}
	if f == 0 {
		h++
	}
	if b == c {
		break
	}
	b += 17
}
*/

func main() {
	h := 0
	for b := 109300; b <= 126300; b += 17 {
		for d := 2; d < b; d++ {
			if b%d == 0 {
				h++
				break
			}
		}
	}

	fmt.Println(h)
}
