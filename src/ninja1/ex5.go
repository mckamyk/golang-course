package main

import "fmt"

type foo int

var x foo

func main() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Printf("%v\t%T\n", x, x)
	y := int(x)
	fmt.Printf("%v\t%T\n", y, y)
}
