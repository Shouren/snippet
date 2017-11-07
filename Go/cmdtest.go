package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := "sleep 20"
	binary := "sh"
	param := "-c"
	fmt.Println(cmd)
	p := exec.Command(binary, param, cmd)
	p.Start()
	fmt.Println("Wait to finish ......")
	p.Wait()
}
