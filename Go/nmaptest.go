package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	var host string
	host = "103.37.149.30"
	fmt.Println("Nmap host:", host)

	out, err := exec.Command("nmap", "-oG", "-", "-sn", "-PA22", host).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
