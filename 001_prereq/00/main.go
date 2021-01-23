package main

import "fmt"

var packageLevelScopeVar int

var s int

func main()  {
	x := 7
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	packageLevelScopeVar = 10
	fmt.Println(packageLevelScopeVar)
	fmt.Println(s)
}
