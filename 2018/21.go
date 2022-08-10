package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/DrJosh9000/exp"
)

// Approach:
// 1. Transpile the program into Go.
// 2. The human then cleans it up and reverse engineers it.
// Hint: it looks like a toy RNG

var tmpl = template.Must(template.New("trans_rights.go").Parse(`package main
	
func run(r0 int) {
	var r1, r2, r3, r4, r5 int
	for {
		switch r{{.IP}} {
		{{range $i, $x := .Program}}case {{$i}}: {{$x}}
		{{end}}default: return
		}
		r{{.IP}}++
	}
}
`))

func relop(op string, a, b, c int) string {
	switch {
	case a == c:
		return fmt.Sprintf("r%d %s= r%d", c, op, b)
	case b == c:
		return fmt.Sprintf("r%d %s= r%d", c, op, a)
	default:
		return fmt.Sprintf("r%d = r%d %s r%d", c, a, op, b)
	}
}

func imop(op string, a, b, c int) string {
	if a == c {
		return fmt.Sprintf("r%d %s= %d", c, op, b)
	} 
	return fmt.Sprintf("r%d = r%d %s %d", c, a, op, b)
}

func main() {
	type program struct {
		IP int
		Program []string
	}
	p := program{}
	exp.MustForEachLineIn("inputs/21.txt", func(line string) {
		if strings.HasPrefix(line, "#ip") {
			if _, err := fmt.Sscanf(line, "#ip %d", &p.IP); err != nil {
				log.Fatalf("Couldn't parse directive: %v", err)
			}
			return
		}
		
		var opcode string
		var a, b, c int
		if _, err := fmt.Sscanf(line, "%s %d %d %d", &opcode, &a, &b, &c); err != nil {
			log.Fatalf("Couldn't parse instruction: %v", err)
		}
		
		var impl string
		switch opcode {
		case "addr":
			impl = relop("+", a, b, c)
		case "addi":
			if a == c {
				if b == 1 {
					impl = fmt.Sprintf("r%d++", c)
					if c == p.IP {
						impl += fmt.Sprintf(" // goto %d", len(p.Program)+2)
					}
				} else {
					impl = fmt.Sprintf("r%d += %d", c, b)
				}
			} else {
				impl = fmt.Sprintf("r%d = r%d + %d", c, a, b)
			}
		case "mulr":
			impl = relop("*", a, b, c)
		case "muli":
			impl = imop("*", a, b, c)
		case "banr":
			impl = relop("&", a, b, c)
		case "bani":
			impl = imop("&", a, b, c)
		case "borr":
			impl = relop("|", a, b, c)
		case "bori":
			impl = imop("|", a, b, c)
		case "setr":
			impl = fmt.Sprintf("r%d = r%d", c, a)
		case "seti":
			impl = fmt.Sprintf("r%d = %d", c, a)
			if c == p.IP {
				impl += fmt.Sprintf(" // goto %d", a+1)
			}
		case "gtir":
			impl = fmt.Sprintf("if %d > r%d { r%d = 1 } else { r%d = 0 }", a, b, c, c)
		case "gtri":
			impl = fmt.Sprintf("if r%d > %d { r%d = 1 } else { r%d = 0 }", a, b, c, c)
		case "gtrr":
			impl = fmt.Sprintf("if r%d > r%d { r%d = 1 } else { r%d = 0 }", a, b, c, c)
		case "eqir":
			impl = fmt.Sprintf("if %d == r%d { r%d = 1 } else { r%d = 0 }", a, b, c, c)
		case "eqri":
			impl = fmt.Sprintf("if r%d == %d { r%d = 1 } else { r%d = 0 }", a, b, c, c)
		case "eqrr":
			impl = fmt.Sprintf("if r%d == r%d { r%d = 1 } else { r%d = 0 }", a, b, c, c)
		}
		p.Program = append(p.Program, impl)
	})
	
	f, err := os.Create("21a-tp.go")
	if err != nil {
		log.Fatalf("Couldn't create file: %v", err)
	}
	defer f.Close()
	if err := tmpl.Execute(f, &p); err != nil {
		log.Fatalf("Couldn't execute template: %v", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("Couldn't close file: %v", err)
	}
	
	log.Print("Transpiled the program...good luck!")
}
