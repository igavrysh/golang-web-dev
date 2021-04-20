package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range get() {
		fmt.Println(n)
		if n == 5 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// gen is a broken generator that will leak a goroutine
func get() <- chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}
