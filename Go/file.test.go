package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("file.test", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte("Test Hello World"))
}
