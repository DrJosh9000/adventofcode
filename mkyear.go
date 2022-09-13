package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("mkyear.go.tmpl"))

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s YEAR", os.Args[0])
	}
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Usage: %s YEAR", os.Args[0])
	}

	if _, err := os.Stat(strconv.Itoa(year)); err == nil {
		log.Fatalf("%d already exists; cowardly refusing to overwrite", year)
	}

	type values struct {
		Y, D int
		P    string
	}

	if err := os.MkdirAll(fmt.Sprintf("%d/inputs", year), 0755); err != nil {
		log.Fatalf("Couldn't create directories: %v", err)
	}

	parts := []string{"a", "b"}
	for d := 1; d <= 25; d++ {
		if d == 25 {
			parts = parts[:1]
		}
		for _, p := range parts {
			f, err := os.Create(fmt.Sprintf("%d/%d%s.go", year, d, p))
			if err != nil {
				log.Fatalf("Couldn't create file: %v", err)
			}
			// not deferring f.Close() because all the errors fatal out
			if err := tmpl.Execute(f, values{Y: year, D: d, P: p}); err != nil {
				log.Fatalf("Couldn't execute template: %v", err)
			}
			if err := f.Close(); err != nil {
				log.Fatalf("Couldn't close file: %v", err)
			}
		}
	}
}
