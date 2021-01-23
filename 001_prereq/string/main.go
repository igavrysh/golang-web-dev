package main

import (
	"fmt"
)

func main() {
	s := "Hello"
	fmt.Println("1:", s)
	fmt.Println("2:", []byte(s))
	fmt.Println("3:", string([]byte(s)))

	for i, v := range []byte(s) {
		fmt.Println("4:", i, v)
	}

	fmt.Println("5:", []byte(s)[2])
	fmt.Println("6:", string([]byte(s)[2]))

	fmt.Println("7:", string([]byte(s)[3:]))

}

