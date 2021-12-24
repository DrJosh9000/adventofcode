package main

import (
	"fmt"
	"sync"
)
	
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 9; i++ {
		i := i
		wg.Add(1)
		go func() {
			in := make([]int, 0, 14)
			search(append(in, i))
			wg.Done()
		}()
	}
	wg.Wait()
}

func search(in []int) {
	if len(in) == 14 {
		if eval(in) {
			fmt.Println(in)
		}
		return
	}
	for j := 1; j <= 9; j++ {
		search(append(in, j))
	}
}

func eval(in []int) bool {
	w, x, y, z := 0, 0, 0, 0
	w = in[0]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 1 == 0 {
		return false
	}
	z /= 1
	x += 11
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 14
	y *= x
	z += y
	w = in[1]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 1 == 0 {
		return false
	}
	z /= 1
	x += 13
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 8
	y *= x
	z += y
	w = in[2]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 1 == 0 {
		return false
	}
	z /= 1
	x += 11
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 4
	y *= x
	z += y
	w = in[3]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 1 == 0 {
		return false
	}
	z /= 1
	x += 10
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 10
	y *= x
	z += y
	w = in[4]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 26 == 0 {
		return false
	}
	z /= 26
	x += -3
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 14
	y *= x
	z += y
	w = in[5]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 26 == 0 {
		return false
	}
	z /= 26
	x += -4
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 10
	y *= x
	z += y
	w = in[6]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 1 == 0 {
		return false
	}
	z /= 1
	x += 12
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 4
	y *= x
	z += y
	w = in[7]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 26 == 0 {
		return false
	}
	z /= 26
	x += -8
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 14
	y *= x
	z += y
	w = in[8]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 26 == 0 {
		return false
	}
	z /= 26
	x += -3
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 1
	y *= x
	z += y
	w = in[9]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 26 == 0 {
		return false
	}
	z /= 26
	x += -12
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 6
	y *= x
	z += y
	w = in[10]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 1 == 0 {
		return false
	}
	z /= 1
	x += 14
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 0
	y *= x
	z += y
	w = in[11]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 26 == 0 {
		return false
	}
	z /= 26
	x += -6
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 9
	y *= x
	z += y
	w = in[12]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 1 == 0 {
		return false
	}
	z /= 1
	x += 11
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 13
	y *= x
	z += y
	w = in[13]
	x *= 0
	x += z
	if x < 0 || 26 <= 0 {
		return false
	}
	x %= 26
	if 26 == 0 {
		return false
	}
	z /= 26
	x += -12
	if x == w {
		x = 1
	} else { 
		x = 0
	}
	if x == 0 {
		x = 1
	} else { 
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 12
	y *= x
	z += y
	return z == 0
}
