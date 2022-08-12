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
				fmt.Println(r4)
				return
				/*if r4 == r0 {
					return
				}
				break*/
			}
			r3 /= 256
		}
	}
}

func main() {
	run(-1)
}
