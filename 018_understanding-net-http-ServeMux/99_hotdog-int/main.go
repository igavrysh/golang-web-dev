package main

import "fmt"

type hotdog int

func main() {
	var x hotdog
	x = 7

	var y int
	y = int(x)
	fmt.Println(y)
}
