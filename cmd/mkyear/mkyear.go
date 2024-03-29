/*
   Copyright 2022 Josh Deprez

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// The mkyear command creates a new directory containing 49 templated Go source
// files, one for each day and part of an Advent of Code year.
package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

//go:embed mkyear.go.tmpl
var templateSrc string

var tmpl = template.Must(template.New("template").Parse(templateSrc))

func main() {
	year := flag.Int("year", time.Now().Year(), "Generates files for this year")
	fromDay := flag.Int("fromday", 1, "Generates from this day onwards")
	flag.Parse()

	type values struct {
		Y, D int
		P    string
	}

	if err := os.MkdirAll(fmt.Sprintf("%d/inputs", *year), 0755); err != nil {
		log.Fatalf("Couldn't create directories: %v", err)
	}

	parts := []string{"a", "b"}
	for d := *fromDay; d <= 25; d++ {
		parts := parts
		if d == 25 {
			parts = []string{""}
		}
		for _, p := range parts {
			if err := os.MkdirAll(fmt.Sprintf("%d/%d%s", *year, d, p), 0755); err != nil {
				log.Fatalf("Couldn't create directories: %v", err)
			}
			mainpath := fmt.Sprintf("%d/%d%s/main.go", *year, d, p)
			if _, err := os.Stat(mainpath); err == nil {
				log.Fatalf("%s already exists; cowardly refusing to overwrite", mainpath)
			}
			f, err := os.Create(mainpath)
			if err != nil {
				log.Fatalf("Couldn't create file: %v", err)
			}
			// not deferring f.Close() because all the errors fatal out
			if err := tmpl.Execute(f, values{Y: *year, D: d, P: p}); err != nil {
				log.Fatalf("Couldn't execute template: %v", err)
			}
			if err := f.Close(); err != nil {
				log.Fatalf("Couldn't close file: %v", err)
			}
		}
	}
}
