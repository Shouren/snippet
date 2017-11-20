package main

import (
	"fmt"
	"time"
)

func runLoop(Queue chan time.Time) {
	num := 0
	for range Queue {
		if num < 10 {
			tick := <-time.Tick(time.Second)
			Queue <- tick
			num += 1
			fmt.Println(num)
			continue
		}
		return
	}
}

func main() {
	q := make(chan time.Time, 1)
	tick := <-time.Tick(time.Second)
	q <- tick
	runLoop(q)
}
