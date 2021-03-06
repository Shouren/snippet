package main

import (
	"fmt"
)

func main() {
	test := map[string]bool{}

	test["a"] = true
	test["b"] = true

	ret, err := test["b"]
	if !err {
		fmt.Println(ret)
	}
	ret = test["c"]
	if !err {
		fmt.Println(ret)
	}

	fmt.Println("===")
	for k, v := range test {
		fmt.Println(k)
		fmt.Println(v)
	}
	x := test
	x["c"] = false

	fmt.Println("===")
	for k, v := range test {
		fmt.Println(k)
		fmt.Println(v)
	}
}
