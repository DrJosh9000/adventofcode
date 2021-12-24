package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	f, err := os.Open("inputs/24.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}
	defer f.Close()

	tp, err := os.Create("24a-tp.go")
	if err != nil {
		log.Fatalf("Couldn't open intermediate: %v", err)
	}
	defer tp.Close()

	_, err = tp.WriteString(`package main

import (
	"fmt"
	"sync"
)
	
func main() {
	fmt.Println("starting search")
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
`)
	if err != nil {
		log.Fatalf("Couldn't write preamble: %v", err)
	}

	nextvar := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		token := strings.Split(sc.Text(), " ")
		if token[0] == "inp" {
			fmt.Fprintf(tp, "\t%s = in[%d]\n", token[1], nextvar)
			nextvar++
			continue
		}
		switch token[0] {
		case "add":
			fmt.Fprintf(tp, "\t%s += %s\n", token[1], token[2])
		case "mul":
			fmt.Fprintf(tp, "\t%s *= %s\n", token[1], token[2])
		case "div":
			fmt.Fprintf(tp, `	if %s == 0 {
		return false
	}
	%s /= %s
`, token[2], token[1], token[2])
		case "mod":
			fmt.Fprintf(tp, `	if %s < 0 || %s <= 0 {
		return false
	}
	%s %%= %s
`, token[1], token[2], token[1], token[2])
		case "eql":
			fmt.Fprintf(tp, `	if %s == %s {
		%s = 1
	} else { 
		%s = 0
	}
`, token[1], token[2], token[1], token[1])
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}

	tp.WriteString(`	return z == 0
}
`)
	if err := tp.Close(); err != nil {
		log.Fatalf("Couldn't close intermediate: %v", err)
	}

	cmd := exec.Command("go", "run", "24a-tp.go")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalf("Couldn't run intermediate: %v", err)
	}
}
