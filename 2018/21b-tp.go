package main

import "fmt"
	
func run(r0 int) {
	var r3, r4 int
	
	r4 = 123
	for {
		r4 &= 456
		if r4 == 72 { 
			break
		}
	}
	
	seen := make(map[int]struct{})
	prev := -1
	
	r4 = 0
	for {
		r3 = r4 | 65536
		r4 = 16098955
		
		for {
			r4 += r3 & 255
			r4 &= 16777215
			r4 *= 65899
			r4 &= 16777215
			if 256 > r3 {
				if _, s := seen[r4]; s {
					fmt.Println(prev)
					return
				}
				seen[r4] = struct{}{}
				prev = r4
				
				if r4 == r0 {
					return
				}
				break
			}
			r3 /= 256
		}
	}
}

func main() {
	run(-1)
}
