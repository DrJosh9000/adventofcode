package main

import "fmt"

func main() {
	const N = 1000000
	succ := []int{
		5: 9,
		9: 8,
		8: 1,
		1: 6,
		6: 2,
		2: 7,
		7: 3,
		3: 4,
		4: 10,
		N: 5,
	}
	for i := 10; i < N; i++ {
		succ[i] = i + 1
	}

	cur := 5
	for i := 0; i < 10000000; i++ {
		dest := cur
		for {
			dest = (dest+N-2)%N + 1
			if dest == succ[cur] {
				continue
			}
			if dest == succ[succ[cur]] {
				continue
			}
			if dest != succ[succ[succ[cur]]] {
				break
			}
		}
		succ[dest], succ[cur], succ[succ[succ[succ[cur]]]] = succ[cur], succ[succ[succ[succ[cur]]]], succ[dest]
		cur = succ[cur]
	}
	fmt.Println(succ[1] * succ[succ[1]])
}
