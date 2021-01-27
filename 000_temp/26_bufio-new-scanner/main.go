package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "I felt so good like anything was\nMynameisthenya"
	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		fmt.Println("next scan")
		line := scanner.Text()
		fmt.Println(line)
	}
}