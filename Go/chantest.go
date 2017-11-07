package main

import (
	"fmt"
	"time"
)

func runLoop(Queue <-chan time.Time) {
	num := 0
	for range Queue {
		if num < 10 {
			Queue <- time.Tick(time.Second)
			num = +1
		}
		fmt.Println(num)
	}
}

func main() {
	runLoop(time.Tick(time.Second))
}
