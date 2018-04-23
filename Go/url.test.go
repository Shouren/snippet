package main

import (
	"fmt"
	"net/url"
)

func main() {
	//content := "ftp://127.0.0.1:22/abc/name.txt"
	content := "file:///data/deploy.yaml"

	u, err := url.Parse(content)

	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
}
