package main

import (
	"fmt"
)

func main() {
	// Capacity with 2
	array := make([]int, 0, 2)
	array = append(array, 1)
	a := array
	b := array
	a = append(a, 2)
	b = append(b, 3)
	fmt.Println("Slice with capacity 2")
	fmt.Println("O: ", array)
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)

	// Capacity not specified
	array = make([]int, 0)
	array = append(array, 1)
	c := array
	d := array
	c = append(c, 2)
	d = append(d, 3)
	fmt.Println("Slice with capacity not specified")
	fmt.Println("O: ", array)
	fmt.Println("C: ", c)
	fmt.Println("D: ", d)
}
