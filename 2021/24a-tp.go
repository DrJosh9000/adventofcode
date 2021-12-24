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
	x, z := 0, 0
	z += in[0] + 14
	z *= 26
	z += in[1] + 8
	z *= 26
	z += in[2] + 4
	z *= 26
	z += in[3] + 10
	x = z%26 - 3
	z /= 26
	if x == in[4] {
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1
	z += (in[4] + 14) * x
	x = z%26 - 4
	z /= 26
	if x == in[5] {
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1
	z += (in[5] + 10) * x
	z *= 26
	z += in[6] + 4
	x = z%26 - 8
	z /= 26
	if x == in[7] {
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1
	z += (in[7] + 14) * x
	x = z%26 - 3
	z /= 26
	if x == in[8] {
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1
	z += (in[8] + 1) * x
	x = z%26 - 12
	z /= 26
	if x == in[9] {
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1
	z += (in[9] + 6) * x
	z *= 26
	z += in[10]
	x = z%26 - 6
	z /= 26
	if x == in[11] {
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1
	z += (in[11] + 9) * x
	z *= 26
	z += in[12] + 13
	x = z%26 - 12
	z /= 26
	if x == in[13] {
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1
	z += (in[13] + 12) * x
	atomic.AddUint64(&evals, 1)
	return z == 0
}
