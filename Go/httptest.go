package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var url string
	url = "http://www.google.com/robots.txt"
	fmt.Printf("Getting from %s ...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(ret)
}
