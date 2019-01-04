package main

import (
	"flag"
	"fmt"
)

var skiilParam = flag.String("skill", "", "skill to perform")

func main() {
	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skiilParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
