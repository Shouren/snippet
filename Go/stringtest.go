package main

import (
	"fmt"
	"strings"
)

func main() {
	// Connect
	a := "Hello"
	a = a + " World"
	fmt.Println(a)

	// Split
	b := "hello;world"
	for idx, item := range strings.Split(b, ";") {
		if item != "" {
			fmt.Println(idx, item)
		}
	}

	// Replace
	c := "hello/world:"
	r := strings.NewReplacer("/", "-", ":", "-")
	d := r.Replace(c)
	fmt.Println(c)
	fmt.Println(d)
}
