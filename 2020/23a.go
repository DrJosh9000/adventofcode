package main

import "fmt"

func main() {
	//circ := []int{3,8,9,1,2,5,4,6,7}
	circ := []int{5, 9, 8, 1, 6, 2, 7, 3, 4}
	for i := 0; i < 100; i++ {
		destn := circ[0]
		var dest int
	findDest:
		for {
			destn = (destn+7)%9 + 1
			//fmt.Printf("finding index of %d in %v\n", destn, circ[4:])
			for j, n := range circ[4:] {
				if n == destn {
					dest = j + 4
					break findDest
				}
			}
			//fmt.Println("not found")
		}
		//fmt.Printf("dest is %d\n", dest)
		var c2 []int
		c2 = append(c2, circ[4:dest+1]...)
		c2 = append(c2, circ[1:4]...)
		c2 = append(c2, circ[dest+1:]...)
		c2 = append(c2, circ[0])
		//fmt.Printf("c2 is %v\n", c2)
		circ = c2
		//fmt.Println(circ)
	}

	// Now find 1
	for j, n := range circ {
		if n == 1 {
			for _, x := range circ[j+1:] {
				fmt.Printf("%d", x)
			}
			for _, x := range circ[:j] {
				fmt.Printf("%d", x)
			}
			fmt.Println()
			return
		}
	}

}
