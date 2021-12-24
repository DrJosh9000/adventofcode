package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var evals uint64

func main() {
	fmt.Println("starting search")
	start := time.Now()
	go func() {
		for range time.Tick(time.Minute) {
			e := atomic.LoadUint64(&evals)
			fmt.Println("evals/sec:", float64(e)/time.Since(start).Seconds())
		}
	}()
	var wg sync.WaitGroup
	for i := 9; i >= 1; i-- {
		i := i
		for j := 9; j >= 1; j-- {
			j := j
			wg.Add(1)
			go func() {
				search(append(make([]int, 0, 14), i, j))
				wg.Done()
			}()
		}
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
	for i := 9; i >= 1; i-- {
		search(append(in, i))
	}
}

func eval(in []int) bool {
	w, x, y, z := 0, 0, 0, 0
	z = (in[0] + 14) * 26
	z += (in[1] + 8) * 26
	z += (in[2] + 4) * 26
	x = 1
	z += in[3] + 10
	w = in[4]
	x = z % 26
	z /= 26
	x += -3
	if x == in[4] {
		x = 0
	} else {
		x = 1
	}
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 14
	y *= x
	z += y
	w = in[5]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 10
	y *= x
	z += y
	w = in[6]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 4
	y *= x
	z += y
	w = in[7]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 14
	y *= x
	z += y
	w = in[8]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y++
	y *= x
	z += y
	w = in[9]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 6
	y *= x
	z += y
	w = in[10]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 0
	y *= x
	z += y
	w = in[11]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 9
	y *= x
	z += y
	w = in[12]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 13
	y *= x
	z += y
	w = in[13]
	x = 0
	x += z
	if x < 0 {
		return false
	}
	x %= 26
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
	y = 0
	y += 25
	y *= x
	y++
	z *= y
	y = 0
	y += w
	y += 12
	y *= x
	z += y
	atomic.AddUint64(&evals, 1)
	return z == 0
}
